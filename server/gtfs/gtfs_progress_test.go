package gtfs

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init(){
	fmt.Println("Running gtfs_progress_tests")
}


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

	gtfs := NewGTFSFromZipBytesWithProgress("GTFS.zip",zbytes, messages, errorChannel)
	
	assert.Len(t, gtfs.Agencies, 40)
	assert.Len(t, gtfs.Attributions, 1)
	assert.Len(t, gtfs.CalendarDates, 98)
	assert.Len(t, gtfs.Calendar, 98)
	assert.Len(t, gtfs.Frequencies, 0)
	assert.Len(t, gtfs.Routes, 98)
	assert.Len(t, gtfs.Shapes, 4)
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

	NewGTFSFromZipBytesWithProgress("GTFS.zip",zbytes, messages, errorChannel)
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

	NewGTFSFromZipBytesWithProgress("GTFS.zip",zbytes, messages, errorChannel)
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

	NewGTFSFromZipBytesWithProgress("GTFS.zip",zbytes, messages, errorChannel)
	assert.ErrorContains(t, err,"Error: 'agency.txt' missing required field(s)")
}
