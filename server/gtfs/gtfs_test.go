package gtfs

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"io"
	"os"
	"bufio"

)

func init(){
	fmt.Println("Running gtfs_tests")
}


// Testing based on smaller GTFSDK.zip file. 
func TestParseGTFSZipIntoGTFSFiles(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}


	gtfsFiles, err := ParseGTFSZipIntoGTFSFiles(zbytes)
	if err != nil {
		t.Error("Error ParseGTFSZipIntoGTFSFiles!")
	}

	assert.Equal(t, "agency.txt", gtfsFiles[0].Name)
	assert.Equal(t, "attributions.txt", gtfsFiles[1].Name)
	assert.Equal(t, "calendar.txt", gtfsFiles[2].Name)
	assert.Equal(t, "calendar_dates.txt", gtfsFiles[3].Name)
	assert.Equal(t, "frequencies.txt", gtfsFiles[4].Name)
	assert.Equal(t, "routes.txt", gtfsFiles[5].Name)
	assert.Equal(t, "shapes.txt", gtfsFiles[6].Name)
	assert.Equal(t, "stops.txt", gtfsFiles[7].Name)
	assert.Equal(t, "stop_times.txt", gtfsFiles[8].Name)
	assert.Equal(t, "transfers.txt", gtfsFiles[9].Name)
	assert.Equal(t, "trips.txt", gtfsFiles[10].Name)
}

func TestMarshalGTFSAgency(t *testing.T){
	
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}


	gtfsFiles, err := ParseGTFSZipIntoGTFSFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalGTFSAgency!")
	}

	data := gtfsFiles[0]
	agencies, err := UnmarshallAgencies(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalGTFSAgency: " + err.Error())
	}

	assert.Len(t, agencies, 40)
}


func TestMarshalGTFSAttribution(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseGTFSZipIntoGTFSFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalGTFSAttribution!")
	}

	data := gtfsFiles[1]
	attributions, err := UnmarshallAttributions(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalGTFSAttribution: " + err.Error())
	}
	assert.Len(t, attributions, 1)
}


func TestMarshalGTFSCalendar(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseGTFSZipIntoGTFSFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalGTFSCalendar!")
	}

	data := gtfsFiles[2]
	calendar, err := UnmarshallCalendar(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalGTFSCalendar: " + err.Error())
	}
	assert.Len(t, calendar, 98)
}


func TestMarshalGTFSCalendarDate(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseGTFSZipIntoGTFSFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalGTFSCalendarDates!")
	}

	data := gtfsFiles[3]
	calendar, err := UnmarshallCalendarDate(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalGTFSCalendarDates: " + err.Error())
	}
	assert.Len(t, calendar, 98)
}


func TestMarshalGTFSFrequency(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseGTFSZipIntoGTFSFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalGTFSFrequency!")
	}

	data := gtfsFiles[4]
	frequency , err := UnmarshallFrequency(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalGTFSFrequency: " + err.Error())
	}
	assert.Len(t, frequency, 0)
}


func TestMarshalGTFSRoute(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseGTFSZipIntoGTFSFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalGTFSRoute!")
	}

	data := gtfsFiles[5]
	routes, err := UnmarshallRoutes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalGTFSRoute: " + err.Error())
	}
	assert.Len(t, routes, 98)
}



func TestMarshalGTFSShape(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseGTFSZipIntoGTFSFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalGTFSShape!")
	}

	data := gtfsFiles[6]
	shapes, err := UnmarshallShapes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalGTFSShape: " + err.Error())
	}
	
	assert.Len(t, shapes, 4)
}

/*
// Testing based on smaller GTFSDK.zip file. 
func TestUnzipGTFSFromBytes(t *testing.T){
	
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}
	messages := make(chan GTFSLoadProgress)
	errorChannel := make(chan error)
	
	go func(){
		for {
			select {
				case msg := <-messages:
					if msg.Filename == "GTFS.zip" && msg.Done {
						return
					}
			
				case err := <-errorChannel:
					t.Error("Unexpected Error: " + err.Error())
					return
			}
		}
	}()

	gtfs := NewGTFSFromZipBytes("GTFS.zip",zbytes, messages, errorChannel)
	
	assert.Len(t, gtfs.Agencies, 40)
	assert.Len(t, gtfs.Attributions, 1)
	assert.Len(t, gtfs.CalendarDates, 98)
	assert.Len(t, gtfs.Calendar, 98)
	assert.Len(t, gtfs.Frequencies, 0)
	assert.Len(t, gtfs.Routes, 98)
	assert.Len(t, gtfs.Shapes, 98)
	assert.Len(t, gtfs.StopTimes, 98)
	assert.Len(t, gtfs.Stops, 98)
	assert.Len(t, gtfs.Transfers, 98)
	assert.Len(t, gtfs.Trips, 98)
}

func TestUnzipGTFSFromBytesTooManyColumns(t *testing.T){
	
	zbytes, err := getBytesFromZipFile("test_data/GTFS_TOO_MANY_AGENCY_COLUMNS.zip")
	if err != nil {
		t.Error(err.Error())
	}
	messages := make(chan GTFSLoadProgress)
	errorChannel := make(chan error)
	
	go func(){
		for {
			select {
				case msg := <-messages:
					if msg.Filename == "GTFS.zip" && msg.Done {
						return
					}
			
				case err = <-errorChannel:
					break
			}
		}
	}()

	NewGTFSFromZipBytes("GTFS.zip",zbytes, messages, errorChannel)
	assert.ErrorContains(t, err,"Error reading file 'agency.txt'")
}

func TestUnzipGTFSFromBytesMissingColumn(t *testing.T){
	
	zbytes, err := getBytesFromZipFile("test_data/GTFS_MISSING_AGENCY_COLUMN.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}
	messages := make(chan GTFSLoadProgress)
	errorChannel := make(chan error)
	
	go func(){
		for {
			select {
				case msg := <-messages:
					if msg.Filename == "GTFS.zip" && msg.Done {
						return
					}
			
				case err = <-errorChannel:
					break
			}
		}
	}()

	NewGTFSFromZipBytes("GTFS.zip",zbytes, messages, errorChannel)
	assert.ErrorContains(t, err,"Error reading file 'agency.txt'")
}

func TestUnzipGTFSFromBytesWrongAgencyField(t *testing.T){
	
	zbytes, err := getBytesFromZipFile("test_data/GTFS_WRONG_AGENCY_FIELD.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}
	messages := make(chan GTFSLoadProgress)
	errorChannel := make(chan error)
	
	go func(){
		for {
			select {
				case msg := <-messages:
					if msg.Filename == "GTFS.zip" && msg.Done {
						return
					}
			
				case err = <-errorChannel:
					break
			}
		}
	}()

	NewGTFSFromZipBytes("GTFS.zip",zbytes, messages, errorChannel)
	assert.ErrorContains(t, err,"Error: 'agency.txt' missing required field(s)")
}
*/

func getBytesFromZipFile(path string) (zbytes []byte, err error) {
	
	gtfs, err := filepath.Abs(path)
	if err != nil {
		return
	}

	file, err := os.Open(gtfs)
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat();
	if err != nil {
		return
	}

	zbytes = make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(zbytes)
	if err != nil && err != io.EOF {
		return
	}

	return
}
