package gtfs

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"encoding/csv"
	"strings"
	"reflect"
	"math"
	"time"
	"net/http"
)

type GTFSCSVFile struct {
	Header []string
	Records [][]string
	Name string
}


func DownloadGTFS(url string) (bytes []byte, err error){
	
	n := time.Now()
	// need to be 24 hours when live
	day, err := time.ParseDuration("24h")
	if err != nil {
		return
	}
	ts := n.Add(-day).Format(time.RFC1123)

	client := &http.Client{
		Timeout:10*time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Add("If-Modified-Since", ts)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// if not 200 or 304, we need to log this as failed. E.g. resource moved or something like that.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		err = fmt.Errorf("Error downloading GTFS zip file from '%s'. Status: '%s'", url, resp.Status)
		return 
	}

	bytes, err = io.ReadAll(resp.Body)
	
	return 
}

// we need a test for this as well.
func ParseGTFSZipIntoGTFSFiles (zbytes []byte) (files []GTFSCSVFile, err error){
	reader := bytes.NewReader(zbytes)
    zreader, err := zip.NewReader(reader, int64(len(zbytes)))
	if err != nil {
		return
	}

	for _, file := range zreader.File {
		freader, err := file.Open()
		if err != nil {
			break
		}

		defer func(){
			if err = freader.Close(); err != nil {
				panic(err) // this should panic the process
			  }
		}()

		r := csv.NewReader(freader) // this should be the entry for something else
		data, err := r.ReadAll()
		if err != nil {
			break
		}

		files = append(files, GTFSCSVFile{data[0], data[1:], file.Name})
	}

	// we get zip files but not real files. 	

	return 
}


// Modified from by https://github.com/artonge/go-csv-tag/blob/4b40f225e91a009021bac2ae6fd04a3d90c58b12/load.go#L142
// Unmarshals the rows from the zipped GTFS (csv like) files.
// takes filename for progress reporting
// takes a csv header string slice
// takes a row of string slices with the csv valuies
// destination interface to unmarshal into
// message channel to report progress
// error channel for error reporting
func UnmarshalSlice(filename string, header []string, rows[][]string, destination interface{}) (err error){
	
	// developer error
	if destination == nil {
		err = fmt.Errorf("Error Unmarshalling: Destination slice is nil")
		return
	}

	// developer error
	if reflect.TypeOf(destination).Elem().Kind() != reflect.Slice { 
		err = fmt.Errorf("Error Unmarshalling: Destination is not a slice")
		return
	}
	
	// Map each header name to its index.
	headerIndex := make(map[string]int)
	for i, name := range header {
		headerIndex[strings.TrimSpace(name)] = i
	}

	// Create the slice to put the values in.
	refSlice := reflect.MakeSlice(
		reflect.ValueOf(destination).Elem().Type(),
		len(rows),
		len(rows),
	)

	if ok := hasRequiredFields(headerIndex, refSlice.Index(0)); !ok {
		err = fmt.Errorf("Error: '%s' missing required field(s)", filename)
		return
	}

	percent := 0.00

	for i, row := range rows {
		refStruct := refSlice.Index(i)

		status := math.Floor(float64(i) / float64(len(rows)) * 100.00)
		if status > percent+1 { //we just want a few less status messages
			percent = status
		}
		
		n := refStruct.NumField()
		for j := 0; j < n ; j++ {
			csvTag := refStruct.Type().Field(j).Tag.Get("csv")
			if csvTag == "" {
				continue
			}

			position, ok := headerIndex[csvTag] 
			if !ok {
				continue
			}

			err := storeValue(row[position], refStruct.Field(j))
			if err != nil {
				err = fmt.Errorf("line: %v to slice: %v:\n	==> %v", row, refStruct, err)
				break
			}
		}
		
		if err != nil {
			break
		}
	}

	reflect.ValueOf(destination).Elem().Set(refSlice)
	return 
}