package gtfs

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
)

func init(){
	fmt.Println("Running repository_tests")
}

func TestIngestStops(t *testing.T){

	godotenv.Load("sql/database.env")

	
	zbytes, err := getBytesFromZipFile("test_data/GTFSDK.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}

	gtfsFiles, err := ParseZipIntoFiles(zbytes)
	if err != nil {
		t.Error("Error TestIngestStops!")
	}

	data := gtfsFiles[7]
	stops, err := UnmarshallStops(data.Header, data.Records)
	if err != nil {
		t.Error("Error TestIngestStops: " + err.Error())
	}

	repo, err := NewRepository()
	if err != nil {
		t.Error("Error TestIngestStops: " + err.Error())
	}

	err = repo.IngestStops(stops)
	assert.NoError(t, err)
}