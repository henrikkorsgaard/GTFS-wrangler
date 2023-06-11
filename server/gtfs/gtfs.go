package gtfs

import (
	"archive/zip"
	"bytes"
	"sync"
	"io"
	"encoding/csv"
	"fmt"
	"strings"
	"reflect"
	"strconv"
	"math"
)

// Do we need the filename of the original zip?
// Just to be able to indicate to the user, that the 
// full zip is processed?
func NewGTFSFromZipBytes(filename string, zbytes []byte, messenges chan GTFSLoadProgress, errorChannel chan error) (gtfs GTFS) {

	reader := bytes.NewReader(zbytes)
    zreader, err := zip.NewReader(reader, int64(len(zbytes)))
	if err != nil {
		return
	}
	/*
	The recommended way is to open a file and then get a ReadCloser, to close the zip file.
	See: https://pkg.go.dev/archive/zip#ReadCloser.Close

	But since we are not providing any zipfile, ReadCloser.Close() will fail, because it is trying to close a non-existing file. See
	https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/archive/zip/reader.go;drc=145dd38471fe5e14b8a77f5f466b70ab49c9a62b;l=51

	Conclusion, we do not use a closer, because there is no file to close.
	*/
	//https://gobyexample.com/waitgroups
	var wg sync.WaitGroup

	for _, file := range zreader.File {

		freader, err := file.Open()
		if err != nil {
			// we return error and end for now
			// if there are non-consequential errors
			// then we can filter these and skip the file
			// with continue
			errorChannel<-err
			return
		}
		defer func() { // we need to handle the potential error from close
			if err = freader.Close(); err != nil {
			  panic(err) // this should panic the process
			}
		}()// call the anon defer func
		
		// do in a nother function
		var destination any

		switch file.Name {
			case "agency.txt":
				destination = &gtfs.Agencies
			case "stops.txt":
				destination = &gtfs.Stops
			case "routes.txt":
				destination = &gtfs.Routes
			case "trips.txt":
				destination = &gtfs.Trips
			case "stop_times.txt":
				destination = &gtfs.StopTimes
			case "calendar.txt":
				destination = &gtfs.Calendar
			case "calendar_dates.txt":
				destination = &gtfs.CalendarDates
			case "fare_attributes.txt":
				destination = &gtfs.FareAttributes
			case "fare_rules.txt":
				destination = &gtfs.FareRules
			case "shapes.txt":
				destination = &gtfs.Shapes
			case "frequencies.txt":
				destination = &gtfs.Frequencies
			case "transfers.txt":
				destination = &gtfs.Transfers
			case "pathways.txt":
				destination = &gtfs.Pathways
			case "levels.txt":
				destination = &gtfs.Levels
			case "feed_info.txt":
				destination = &gtfs.FeedInfo
			case "translations.txt":
				destination = &gtfs.Translations
			case "attributions.txt":
				destination = &gtfs.Attributions
			default:
				continue
		}
		
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			loadGTFSFileFromByteReader(filename, freader,destination,messenges, errorChannel)
			return
		}(file.Name)
	}
	
	wg.Wait()
	
	messenges<-GTFSLoadProgress{FileName: filename, Message:"Done loading all GTFS files", Done: true}
	return gtfs
}

func loadGTFSFileFromByteReader(filename string, freader io.Reader, destination interface{}, messages chan GTFSLoadProgress, errorChannel chan error) {
	r := csv.NewReader(freader) // this should be the entry for something else
	data, err := r.ReadAll()
	if err != nil {
		errorChannel <- err
		return
	}

	if len(data) < 2 {
		messages <- GTFSLoadProgress{filename,100,1,1,"file only contained header",true}
		return
	}

	header := data[0]
	rows := data[1:] 
	
	unmarshalSlice(filename, header, rows, destination, messages, errorChannel)
}


// Modified from by https://github.com/artonge/go-csv-tag/blob/4b40f225e91a009021bac2ae6fd04a3d90c58b12/load.go#L142
func unmarshalSlice(filename string, header []string, rows[][]string, destination interface{}, messages chan GTFSLoadProgress, errorChannel chan error) {
	
	if destination == nil {
		errorChannel<-fmt.Errorf("Error Unmarshalling: Destination slice is nil")
		return
	}

	if reflect.TypeOf(destination).Elem().Kind() != reflect.Slice { 
		errorChannel<-fmt.Errorf("Error Unmarshalling: Destination is not a slice")
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
		errorChannel<-fmt.Errorf("Error loading '%s' from file: CSV Rows are missing required fields", filename)
		return
	}

	percent := 0.00

	for i, row := range rows {
		refStruct := refSlice.Index(i)

		status := math.Floor(float64(i) / float64(len(rows)) * 100.00)
		if status > percent+1 { //we just want a few less status messages
			percent = status
			messages <- GTFSLoadProgress{filename,int(percent), len(rows),i,"loading",false}
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
				errorChannel<-fmt.Errorf("line: %v to slice: %v:\n	==> %v", row, refStruct, err)
				return
			}
		}
	}

	messages <- GTFSLoadProgress{filename,100, len(rows),len(rows),"completed loading",true}

	reflect.ValueOf(destination).Elem().Set(refSlice)
	return 
}

// Store value
// Ref: https://github.com/artonge/go-csv-tag/blob/4b40f225e91a009021bac2ae6fd04a3d90c58b12/load.go#L194
func storeValue(rawValue string, valRv reflect.Value) error {
	rawValue = strings.TrimSpace(rawValue)
	switch valRv.Kind() {
	case reflect.String:
		valRv.SetString(rawValue)
	case reflect.Int64:
		fallthrough
	case reflect.Int:
		value, err := strconv.ParseInt(rawValue, 10, 64)
		if err != nil && rawValue != "" {
			return fmt.Errorf("error parsing int '%v':\n	==> %v", rawValue, err)
		}
		valRv.SetInt(value)
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil && rawValue != "" {
			return fmt.Errorf("error parsing float '%v':\n	==> %v", rawValue, err)
		}
		valRv.SetFloat(value)
	case reflect.Bool:
		value, err := strconv.ParseBool(rawValue)
		if err != nil && rawValue != "" {
			return fmt.Errorf("error parsing bool '%v':\n	==> %v", rawValue, err)
		}
		valRv.SetBool(value)
	}

	return nil
}

func hasRequiredFields(headerIndex map[string]int, refStruct reflect.Value) bool {
	n := refStruct.NumField()
	for i := 0; i < n ; i++ {
		csvTag := refStruct.Type().Field(i).Tag.Get("csv")
		requiredTag := refStruct.Type().Field(i).Tag.Get("required")
		if csvTag == "" {
			continue
		}

		_, ok := headerIndex[csvTag] //this does not yet account for required!
		if !ok && requiredTag == "true" {
			return false
		}
	}
	return true
}
