package gtfs

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"io"
	"os"
	"bufio"
)


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
					if msg.FileName == "GTFS.zip" && msg.Done {
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
					if msg.FileName == "GTFS.zip" && msg.Done {
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
					if msg.FileName == "GTFS.zip" && msg.Done {
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
					if msg.FileName == "GTFS.zip" && msg.Done {
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

// I need to figure out how to test this and if we need to implement a file check
func TestUnzipGTFSFromBytesMissingAgencyFile(t *testing.T){
	
	zbytes, err := getBytesFromZipFile("test_data/GTFS_MISSING_AGENCY_FILE.zip")
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}
	messages := make(chan GTFSLoadProgress)
	errorChannel := make(chan error)

	go func(){
		for {
			select {
				case msg := <-messages:
					if msg.FileName == "GTFS.zip" && msg.Done {
						return
					}
			
				case err = <-errorChannel:
					break
			}
		}
	}()

	NewGTFSFromZipBytes("GTFS.zip",zbytes, messages, errorChannel)
	assert.ErrorContains(t, err,"Missing 'agency.txt' file")
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
