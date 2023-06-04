package gtfs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestLoadingStopTimeFromCSVFile(t *testing.T){

	progress := make(chan int)
	errs := make(chan error)

	stopTimes := []StopTime{}

	go LoadGTFSFromCSVFilePath("./test_data/stop_times_short.txt", &stopTimes, progress, errs)

	for {
		done := false
		select {
			case err := <- errs:
				t.Error("LoadGTFSFromCSVFilePath returned unexpected error:\n\t ==>" + err.Error())
				done = true
			case p := <- progress:
				if p == -1 {
					done = true
					break
				} 
		}
		if done {
			break
		}
	}

	for _, stoptime := range stopTimes {
		assert.NotNil(t, stoptime.TripID)
		assert.NotNil(t, stoptime.Arrival)
		assert.NotNil(t, stoptime.Departure)
		assert.NotNil(t, stoptime.StopID)
		assert.NotNil(t, stoptime.StopSequence)
	}
}


func TestLoadingStopTimeFromCSVFileMissingFile(t *testing.T){
	_, err := loadStopTimesSlice("./test_data/stop_times_non_existing_file.txt")
	assert.Error(t, err)
}

func TestLoadingStopTimeFromCSVFileMissingRequiredField(t *testing.T){
	_, err := loadStopTimesSlice("./test_data/stop_times_short_missing_required.txt")
	assert.Error(t, err)
}
