package ingest

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"encoding/csv"
	"strings"
	"strconv"
	"reflect"
	"time"
	"net/http"

	"github.com/twpayne/go-geom"
	"henrikkorsgaard.dk/gtfs-service/domain"
)

type CSVFile struct {
	Header []string
	Records [][]string
	Name string
}

func Download(url string) (bytes []byte, err error){
	
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
func ParseZipIntoFiles (zbytes []byte) (files []CSVFile, err error){
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

		files = append(files, CSVFile{data[0], data[1:], file.Name})
	}

	// we get zip files but not real files. 	

	return 
}

func UnmarshallAgency(header []string, rows[][]string) (agency []domain.Agency, err error) {
	err = unmarshalSlice(header, rows, &agency)
	return  
}

func UnmarshallAttributions(header []string, rows[][]string) (attributions []domain.Attribution, err error) {
	err = unmarshalSlice(header, rows, &attributions)
	return  
}

func UnmarshallCalendar(header []string, rows[][]string) (calendar []domain.Calendar, err error) {
	err = unmarshalSlice(header, rows, &calendar)
	return  
}

func UnmarshallCalendarDate(header []string, rows[][]string) (calendarDate []domain.CalendarDate, err error) {
	err = unmarshalSlice(header, rows, &calendarDate)
	return  
}

func UnmarshallFrequency(header []string, rows[][]string) (frequencies []domain.Frequency, err error) {
	err = unmarshalSlice(header, rows, &frequencies)
	return  
}

func UnmarshallRoutes(header []string, rows[][]string) (routes []domain.Route, err error) {
	err = unmarshalSlice(header, rows, &routes)
	return  
}

func UnmarshallShapes(header []string, rows[][]string) (shapes []domain.Shape, err error) {
	err = unmarshalSlice(header, rows, &shapes)
	
	shapeMap := make(map[string]domain.Shape)
	coordMap := make(map[string][]geom.Coord)
	for _, s := range shapes {
		coord := geom.Coord{s.Lon, s.Lat}
		
		if _, ok := shapeMap[s.ID]; ok {
			coordMap[s.ID] = append(coordMap[s.ID], coord)
		} else {
			coordMap[s.ID] = []geom.Coord{coord}
			shapeMap[s.ID] = s
		}
	}
	
	shapes = make([]domain.Shape, 0, len(shapeMap))

	for k, s := range shapeMap {

	    ls := geom.NewLineString(geom.XY)
		ls.SetSRID(4326)
		ls.MustSetCoords(coordMap[k])

		s.GeoLineString = *ls
		shapes = append(shapes, s)
	}
	
	return
}

func UnmarshallStops(header []string, rows[][]string) (stops []domain.Stop, err error) {
	err = unmarshalSlice(header, rows, &stops)
	for i, n := 0, len(stops); i < n ; i++ {
		p := geom.NewPoint(geom.XY)
		p.SetSRID(4326)
		// This will panic -- handle when we implement full logging.
		p.MustSetCoords(geom.Coord{stops[i].Lon,stops[i].Lat})
		stops[i].GeoPoint = *p
	}

	return  
}

func UnmarshallStopTimes(header []string, rows[][]string) (stopTimes []domain.StopTime, err error) {
	err = unmarshalSlice(header, rows, &stopTimes)
	return  
}

func UnmarshallTransfers(header []string, rows[][]string) (transfers []domain.Transfer, err error) {
	err = unmarshalSlice(header, rows, &transfers)
	return  
}

func UnmarshallTrips(header []string, rows[][]string) (trips []domain.Trip, err error) {
	err = unmarshalSlice(header, rows, &trips)
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
func unmarshalSlice(header []string, rows[][]string, destination interface{}) (err error){
	
	if len(rows) == 0 {
		return
	}

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
		err = fmt.Errorf("Error: file missing required field(s)")
		return
	}

	for i, row := range rows {
		refStruct := refSlice.Index(i)
		
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

			err = storeValue(row[position], refStruct.Field(j))
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

// Store value
// Ref: https://github.com/artonge/go-csv-tag/blob/4b40f225e91a009021bac2ae6fd04a3d90c58b12/load.go#L194
// Way to store different values into the struct
// for now we only focus on strings (they will be converted to json anyways)
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