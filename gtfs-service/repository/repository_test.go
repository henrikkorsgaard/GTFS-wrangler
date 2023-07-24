package gtfs

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"

	"henrikkorsgaard.dk/gtfs-service/util"
	"henrikkorsgaard.dk/gtfs-service/ingest"
)

var (
	testDataString string = "../test_data/GTFSDK.zip"
)

func init(){
	fmt.Println("Running repository_tests")
	godotenv.Load("sql/database.env")
}

func TestIngestStops(t *testing.T){
	
	zbytes, err := util.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestStops!")
	}

	data := gtfsFiles[7]
	stops, err := ingest.UnmarshallStops(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestStops: " + err.Error())
	}

	// Singleton, so we will get the same each time anyways!
	repo, err := NewRepository()
	if err != nil {
		t.Error("Error TestIngestStops: " + err.Error())
	}

	err = repo.IngestStops(stops)
	assert.NoError(t, err)
}

func TestIngestRoutes(t *testing.T){

	zbytes, err := util.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestRoutes!")
	}

	data := gtfsFiles[5]
	routes, err := ingest.UnmarshallRoutes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestRoutes: " + err.Error())
	}

	repo, err := NewRepository()
	if err != nil {
		t.Error("Error TestIngestRoutes: " + err.Error())
	}

	err = repo.IngestRoutes(routes)
	assert.NoError(t, err)
	// we need to assert the success as well -- no errors is not enough?
	// this means wiping the database before running each test!
}


func TestIngestTrips(t *testing.T){
	
	zbytes, err := util.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestTrips!")
	}

	data := gtfsFiles[10]
	trips, err := ingest.UnmarshallTrips(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestTrips: " + err.Error())
	}

	repo, err := NewRepository()
	if err != nil {
		t.Error("Error TestIngestTrips: " + err.Error())
	}

	err = repo.IngestTrips(trips)
	assert.NoError(t, err)
}

func TestIngestShapes(t *testing.T){
	
	zbytes, err := util.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestShapes!")
	}

	data := gtfsFiles[6]
	shapes, err := ingest.UnmarshallShapes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestShapes: " + err.Error())
	}

	repo, err := NewRepository()
	if err != nil {
		t.Error("Error TestIngestShapes: " + err.Error())
	}

	err = repo.IngestShapes(shapes)
	assert.NoError(t, err)
}

func TestIngestStopTimes(t *testing.T){
	
	zbytes, err := util.GetBytesFromZipFile(testDataString)
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestStopTimes!")
	}

	data := gtfsFiles[8]
	stoptimes, err := ingest.UnmarshallStopTimes(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestStopTimes: " + err.Error())
	}

	repo, err := NewRepository()
	if err != nil {
		t.Error("Error TestStopTimes: " + err.Error())
	}

	err = repo.IngestStopTimes(stoptimes)
	assert.NoError(t, err)
}

