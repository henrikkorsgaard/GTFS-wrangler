package gtfs

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"path/filepath"
	"io"
	"os"
	"fmt"
	"bufio"
)

/*
	TODO: Generate test data for all file types
	TODO: Generate invalid test data 
*/


func TestUnzipGTFSFromBytes(t *testing.T){
	
	zbytes, err := getBytesFromZipFile()
	if err != nil {
		t.Error("Error unzipping bytes from file: " + err.Error())
	}
	messages := make(chan GTFSLoadProgress)
	errorChannel := make(chan error)

	go func(){
		for {
			select {
				case msg := <-messages:
					fmt.Println(msg)
					
					if msg.FileName == "GTFS.zip" && msg.Done {
						return
					}
			
				case err := <-errorChannel:
					fmt.Println(err)
					return
			}
		}
	}()

	gtfs := NewGTFSFromZipBytes("GTFS.zip",zbytes, messages, errorChannel)
	
	assert.Len(t, gtfs.Agencies, 40)
	assert.Len(t, gtfs.Attributions, 1)
	assert.Len(t, gtfs.CalendarDates, 15666)
	assert.Len(t, gtfs.Calendar, 1657)
	assert.Len(t, gtfs.Frequencies, 0)
	assert.Len(t, gtfs.Routes, 1650)
	assert.Len(t, gtfs.Shapes, 4006014)
	assert.Len(t, gtfs.StopTimes, 4085012)
	assert.Len(t, gtfs.Stops, 38495)
	assert.Len(t, gtfs.Transfers, 64344)
	assert.Len(t, gtfs.Trips, 181373)
}


func getBytesFromZipFile() (zbytes []byte, err error) {
	
	gtfs, err := filepath.Abs("../../data/GTFS.zip")
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
