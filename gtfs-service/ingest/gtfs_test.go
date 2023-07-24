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
func TestParseZipIntoFiles(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error ParseZipIntoFiles!")
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

func TestMarshalAgency(t *testing.T){
	
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}


	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalAgency!")
	}

	data := gtfsFiles[0]
	agencies, err := UnmarshallAgencies(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalAgency: " + err.Error())
	}

	assert.Len(t, agencies, 40)
}


func TestMarshalAttribution(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalAttribution!")
	}

	data := gtfsFiles[1]
	attributions, err := UnmarshallAttributions(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalAttribution: " + err.Error())
	}
	assert.Len(t, attributions, 1)
}


func TestMarshalCalendar(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalCalendar!")
	}

	data := gtfsFiles[2]
	calendar, err := UnmarshallCalendar(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalCalendar: " + err.Error())
	}
	assert.Len(t, calendar, 98)
}


func TestMarshalCalendarDate(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalCalendarDates!")
	}

	data := gtfsFiles[3]
	calendar, err := UnmarshallCalendarDate(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalCalendarDates: " + err.Error())
	}
	assert.Len(t, calendar, 98)
}


func TestMarshalFrequency(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalFrequency!")
	}

	data := gtfsFiles[4]
	frequency , err := UnmarshallFrequency(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalFrequency: " + err.Error())
	}
	assert.Len(t, frequency, 0)
}


func TestMarshalRoute(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalRoute!")
	}

	data := gtfsFiles[5]
	routes, err := UnmarshallRoutes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalSRoute: " + err.Error())
	}
	assert.Len(t, routes, 98)
}

func TestMarshalShape(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalShape!")
	}

	data := gtfsFiles[6]
	shapes, err := UnmarshallShapes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalShape: " + err.Error())
	}
	
	assert.Len(t, shapes, 4)
}

func TestMarshalStops(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalStops!")
	}

	data := gtfsFiles[7]
	stops, err := UnmarshallStops(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalStops: " + err.Error())
	}
	
	assert.Len(t, stops, 98)
}

func TestMarshalStopTimes(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalStopTimes!")
	}

	data := gtfsFiles[8]
	stopTimes, err := UnmarshallStopTimes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalStopTimes: " + err.Error())
	}
	
	assert.Len(t, stopTimes, 98)
}

func TestMarshalTransfers(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalTransfers!")
	}

	data := gtfsFiles[9]
	transfers, err := UnmarshallTransfers(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalTransfers: " + err.Error())
	}
	
	assert.Len(t, transfers, 98)
}

func TestMarshalTrips(t *testing.T){
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestMarshalTrips!")
	}

	data := gtfsFiles[10]
	trips, err := UnmarshallTrips(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestMarshalTrips: " + err.Error())
	}
	
	assert.Len(t, trips, 98)
}

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
