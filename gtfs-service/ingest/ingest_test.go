package ingest

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"

	"henrikkorsgaard.dk/gtfs-service/testutils"
)

/* NOTES:
- We ignore fare_attributes for now because they are not in our use-case or in the DK dataset.
*/

var (
	testDataString string = "../testutils/data/GTFSDK.zip"
)

func init(){
	fmt.Println("Running ingest_tests")
}


func TestParseZipIntoFiles(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
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
	
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}


	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[0]
	agencies, err := UnmarshallAgency(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	assert.Len(t, agencies, 40)
}


func TestMarshalAttribution(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[1]
	attributions, err := UnmarshallAttributions(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	assert.Len(t, attributions, 1)
}


func TestMarshalCalendar(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[2]
	calendar, err := UnmarshallCalendar(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	assert.Len(t, calendar, 98)
}


func TestMarshalCalendarDate(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[3]
	calendar, err := UnmarshallCalendarDate(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	assert.Len(t, calendar, 98)
}


func TestMarshalFrequency(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[4]
	frequency , err := UnmarshallFrequency(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	assert.Len(t, frequency, 0)
}


func TestMarshalRoute(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[5]
	routes, err := UnmarshallRoutes(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	assert.Len(t, routes, 98)
}

func TestMarshalShape(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[6]
	shapes, err := UnmarshallShapes(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	
	assert.Len(t, shapes, 4)
}

func TestMarshalStops(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[7]
	stops, err := UnmarshallStops(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	
	assert.Len(t, stops, 98)
}

func TestMarshalStopTimes(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[8]
	stopTimes, err := UnmarshallStopTimes(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	
	assert.Len(t, stopTimes, 98)
}

func TestMarshalTransfers(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[9]
	transfers, err := UnmarshallTransfers(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	
	assert.Len(t, transfers, 98)
}

func TestMarshalTrips(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[10]
	trips, err := UnmarshallTrips(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}
	
	assert.Len(t, trips, 98)
}
