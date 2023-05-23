package gtfs

import (
	"github.com/google/uuid"
	"os"
	"fmt"
	"archive/zip"
	"path/filepath"
	"io"
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

func unzipGTFS(src string) (gtfsDir string, err error){
	
	reader, err := zip.OpenReader(src)
	if err != nil {
	  return
	}

	defer func(){ //When done we want to close. Since reader.close can error, we want to panic on that
	  if err := reader.Close(); err != nil {
		panic(err) 
	  }
	}()//call the anon defer function
  
	id := uuid.New()
	gtfsDir, err = filepath.Abs(tempDir + "GTFS_" + id.String() + "/")
	if err != nil {
		return
	}
	
	err = os.Mkdir(gtfsDir, os.ModePerm)
	if err != nil {
		return
	}

	//Iterates over zip.File (and not os.File)
	for _, file := range reader.File {
		// Errors from here should be logged somewhere.
		extractZipFile(file, gtfsDir)
	}

	return 
}
  
func extractZipFile(file *zip.File, pathDir string) (err error) {
	
	//Guard function to deal with polluted zips
	//We cannot safeguard against engineered zips at this point. 
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

