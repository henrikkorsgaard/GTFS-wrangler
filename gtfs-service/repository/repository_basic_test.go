package repository

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	"henrikkorsgaard.dk/gtfs-service/testutils"
	"henrikkorsgaard.dk/gtfs-service/ingest"
)

/* NOTES:
- We combine the database ingest and fetch test because we have the expected number of items in the original CSV file. Except for Shape (because we combine uniques).
- The gtfsFiles[index] index depends on the test data organization in the zip. If we change test data, the order of the files may change causing multiple tests to fail. 
*/

var (
	testDataString string = "../testutils/data/GTFSDK.zip"
)

func init(){
	fmt.Println("Running repository basic ingest and fetch tests")
	godotenv.Load("../config_dev.env")
	err := testutils.ResetDatabase("./sql/gtfs.sql")
	if err != nil {
		panic(err)
	}
}

func TestIngestFetchAgency(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[0]
	agency, err := ingest.UnmarshallAgency(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	// Singleton, so we will get the same each time anyways!
	repo, err := NewRepository()

	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	err = repo.IngestAgency(agency)
	assert.NoError(t, err)

	dbAgency, err :=  repo.FetchAgency();
	assert.Equal(t,len(data.Records), len(dbAgency))
}

func TestIngestFetchStops(t *testing.T){
	
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[7]
	stops, err := ingest.UnmarshallStops(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	// Singleton, so we will get the same each time anyways!
	repo, err := NewRepository()

	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	err = repo.IngestStops(stops)
	assert.NoError(t, err)

	dbStops, err :=  repo.FetchStops();
	assert.Equal(t,len(data.Records), len(dbStops))

}

func TestIngestFetchRoutes(t *testing.T){

	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[5]
	routes, err := ingest.UnmarshallRoutes(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	repo, err := NewRepository()
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	err = repo.IngestRoutes(routes)
	assert.NoError(t, err)

	dbRoutes, err :=  repo.FetchRoutes();
	assert.Equal(t,len(data.Records), len(dbRoutes))
}

func TestIngestFetchTrips(t *testing.T){
	
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[10]
	trips, err := ingest.UnmarshallTrips(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	repo, err := NewRepository()

	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	err = repo.IngestTrips(trips)
	assert.NoError(t, err)

	dbTrips, err :=  repo.FetchTrips();
	assert.Equal(t,len(data.Records), len(dbTrips))
}

func TestIngestFetchShapes(t *testing.T){
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[6]
	shapes, err := ingest.UnmarshallShapes(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	repo, err := NewRepository()

	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	err = repo.IngestShapes(shapes)
	assert.NoError(t, err)

	dbShapes, err :=  repo.FetchShapes();
	assert.Equal(t,4, len(dbShapes))
}

func TestIngestFetchStopTimes(t *testing.T){
	
	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[8]
	stoptimes, err := ingest.UnmarshallStopTimes(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	repo, err := NewRepository()
	
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	err = repo.IngestStopTimes(stoptimes)
	assert.NoError(t, err)

	dbStopTimes, err :=  repo.FetchStopTimes();
	assert.NoError(t, err)
	assert.Equal(t,len(data.Records), len(dbStopTimes))
}


func TestIngestFetchCalendar(t *testing.T){

	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[2]
	calendar, err := ingest.UnmarshallCalendar(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	repo, err := NewRepository()
	
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	err = repo.IngestCalendars(calendar)
	assert.NoError(t, err)

	dbCalendar, err :=  repo.FetchCalendars();
	assert.NoError(t, err)
	assert.Equal(t,len(data.Records), len(dbCalendar))
}


func TestIngestFetchCalendarDates(t *testing.T){

	zbytes, err := testutils.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	data := gtfsFiles[3]
	calendarDates, err := ingest.UnmarshallCalendarDate(data.Header, data.Records)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	repo, err := NewRepository()
	
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	err = repo.IngestCalendarDates(calendarDates)
	assert.NoError(t, err)

	dbCalendarDates, err :=  repo.FetchCalendarDates();
	assert.NoError(t, err)
	assert.Equal(t,len(data.Records), len(dbCalendarDates))
}
