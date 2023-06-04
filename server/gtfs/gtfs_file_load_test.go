package gtfs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestLoadingStopTimeFromCSVFileSlice(t *testing.T){

	stopTimes, err := loadStopTimesSlice("./test_data/stop_times.txt")
	if err != nil {
		t.Error("loadStopTimesFromFile returned unexpected error: " + err.Error())
	}

	if len(stopTimes) == 0 {
		t.Error("loadRoutesFromFile returned []StopTime with length of zero!")
	}
	
	for _, stoptime := range stopTimes {
		assert.NotNil(t, stoptime.TripID)
		assert.NotNil(t, stoptime.Arrival)
		assert.NotNil(t, stoptime.Departure)
		assert.NotNil(t, stoptime.StopID)
		assert.NotNil(t, stoptime.StopSequence)
	}
}

func TestLoadingStopTimeFromCSVFileSliceMissingRequiredField(t *testing.T){
	_, err := loadStopTimesSlice("./test_data/stop_times_short_missing_required.txt")
	assert.Error(t, err)
}
