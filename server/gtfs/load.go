package gtfs

import (
	"reflect"
	"os"
	"strconv"
	"encoding/csv"
	"strings"
	"fmt"
)


// Modified from by https://github.com/artonge/go-csv-tag/blob/4b40f225e91a009021bac2ae6fd04a3d90c58b12/load.go#L142
func unmarshalSlice(header []string, rows[][]string, destination interface{}, progress chan int, errChan chan error) {

	if destination == nil {
		errChan <- fmt.Errorf("destination slice is nil")
	}

	if reflect.TypeOf(destination).Elem().Kind() != reflect.Slice {
		errChan <- fmt.Errorf("destination is not a slice")
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
		errChan <- fmt.Errorf("CSV Rows are missing required fields")
	}

	for i, row := range rows {

		//This is the channel to report progress
		progress <- i
		
		refStruct := refSlice.Index(i)
				
		n := refStruct.NumField()
		for j := 0; j < n ; j++ {
			csvTag := refStruct.Type().Field(j).Tag.Get("csv")
			if csvTag == "" {
				continue
			}

			position, ok := headerIndex[csvTag] //this does not yet account for required!
			if !ok {
				continue
			}

			err := storeValue(row[position], refStruct.Field(j))
			if err != nil {
				errChan <- fmt.Errorf("line: %v to slice: %v:\n	==> %v", row, refStruct, err)
			}
		}
	}

	reflect.ValueOf(destination).Elem().Set(refSlice)
	progress <- -1
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

// Set the value of the valRv to rawValue.
// @param rawValue: the value, as a string, that we want to store.
// @param valRv: the reflected value where we want to store our value.
// @return an error if one occurs.
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

//This need to be generalized as well
func loadFromCSVFilePath(filepath string) (data [][]string, err error){

	csvfile, err := os.Open(filepath)
	if err != nil {
		return
	} 
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	data, err = r.ReadAll()

	return
}
