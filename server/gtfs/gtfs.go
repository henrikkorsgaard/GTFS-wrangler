package gtfs

import (
	"github.com/google/uuid"
	"os"
	"fmt"
	"archive/zip"
	"path/filepath"
	"io"
	"bytes"
)

var (
	tempDir string = "../temp/"
	//per spec: https://developers.google.com/transit/gtfs/reference#dataset_files
	//map with filenames (I'm using bool for easier lookup, instead of the map[string]struct{} trick. Memory is not an issue now)
	validGTFSfilenames = map[string]bool{"agency.txt":true, "stops.txt":true, "routes.txt":true, "trips.txt":true, "stop_times.txt":true, "calendar.txt":true, "calendar_dates.txt":true, "fare_attributes.txt":true, "fare_rules.txt":true, "shapes.txt":true, "frequencies.txt":true, "transfers.txt":true, "pathways.txt":true, "levels.txt":true,"translations.txt":true,"attributions.txt":true}
)

func init(){
	// A bit crude,but I defer the decision on how to handle GTFS data between sessions a bit
	// This is a wrangler tool after all
	// this is too blunt, I'm missing my gitignore!
	os.RemoveAll(tempDir)
	os.Mkdir(tempDir, os.ModePerm)
}


func UnzipGTFSFromBytes(zbytes []byte) (gtfsDir string, err error) {
	//fast wrap would be to just create the file in the temp dir and then go from there.
	reader := bytes.NewReader(zbytes)
    zreader, err := zip.NewReader(reader, int64(len(zbytes)))
	/*
	The recommended way is to open a file and then get a ReadCloser, to close the zip file.
	See: https://pkg.go.dev/archive/zip#ReadCloser.Close

	But since we are not providing any zipfile, ReadCloser.Close() will fail, because it is trying to close a non-existing file. See
	https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/archive/zip/reader.go;drc=145dd38471fe5e14b8a77f5f466b70ab49c9a62b;l=51

	Conclusion, we do not use a closer, because there is no file to close.
	*/

	gtfsDir, err = createUnzipDir()
	if err != nil {
		return
	}
	
	for _, file := range zreader.File {
		// Errors from here should be logged somewhere.
		extractZipFile(file, gtfsDir)
	}
	
	return
} 
  
func extractZipFile(file *zip.File, pathDir string) (err error) {
	
	//Guard function to deal with polluted zips
	//We cannot safeguard against engineered zips at this point. 
	//but given this will run in a docker, we should be safe. Worst case, the user ruins the docker instance
	if !validGTFSfilenames[file.Name] || file.FileInfo().IsDir() {
		err = fmt.Errorf("Unexpected file/directory in GTFS zip: " + file.Name + ". Skipping unzip of this particular file/directory. See https://developers.google.com/transit/gtfs/reference#dataset_files for expected file names.")
	}

	rc, err := file.Open()
	if err != nil {
	  return 
	}
  
	defer func() { // we need to handle the potential error from close
	  if err = rc.Close(); err != nil {
		panic(err)
	  }
	}()// call the anon defer func
  
	//we want to clear the temp dir first
	path := filepath.Join(pathDir,file.Name)
	if err != nil {
		return
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
	  return
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
  
	_, err = io.Copy(f, rc)
	if err != nil {
	  return 
	}
	return
}


func createUnzipDir() (gtfsDir string, err error) {
	id := uuid.New()
	gtfsDir, err = filepath.Abs(tempDir + "GTFS_" + id.String() + "/")
	if err != nil {
		return
	}
	
	err = os.Mkdir(gtfsDir, os.ModePerm)
	if err != nil {
		return
	}

	return
}
