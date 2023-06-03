package gtfs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// I WILL TEST ALL THE GTFS FILES TO MAKE SURE MY LOAD/UNMARSHAL COVERS THE SPEC
// THIS ALSO HELPS ME UNDERSTAND THE GENERALIZATION

// AGENCY.TXT TESTS
/*
func TestLoadingAgencyFromFileHeaderError(t *testing.T){
	_, err := loadAgencies("./test_data/agency_invalid_header.txt")
	assert.Error(t, err)
}

func TestLoadingAgencyFromFileRowError(t *testing.T){
	_, err := loadAgencies("./test_data/agency_invalid_row.txt")
	assert.Error(t, err)
}

func TestLoadingAgencyFromCSVFile(t *testing.T){

	agencies, err := loadAgencies("./test_data/agency.txt")
	if err != nil {
		t.Error("loadAgenciesFromFile returned unexpected error: " + err.Error())
	}

	if len(agencies) == 0 {
		t.Error("loadAgenciesFromFile returned []Agency with length of zero!")
	}
	
	for _, agency := range agencies {
		assert.NotNil(t, agency.ID)
		assert.NotNil(t, agency.Name)
		assert.NotNil(t, agency.URL)
		assert.NotNil(t, agency.Timezone)
		assert.NotNil(t, agency.Lang)
		assert.NotNil(t, agency.Phone)
	}
}

func TestLoadingRouteFromCSVFile(t *testing.T){

	routes, err := loadRoutes("./test_data/routes.txt")
	if err != nil {
		t.Error("loadRoutesFromFile returned unexpected error: " + err.Error())
	}

	if len(routes) == 0 {
		t.Error("loadRoutesFromFile returned []Agency with length of zero!")
	}
	
	for _, route := range routes {
		assert.NotNil(t, route.ID)
		assert.NotNil(t, route.AgencyID)
		assert.NotNil(t, route.Name)
		assert.NotNil(t, route.LongName)
	}
}
*/
func TestLoadingStopTimeFromCSVFile(t *testing.T){

	stopTimes, err := loadStopTimes("./test_data/stop_times.txt")
	if err != nil {
		t.Error("loadStopTimesFromFile returned unexpected error: " + err.Error())
	}

	if len(stopTimes) == 0 {
		t.Error("loadRoutesFromFile returned []Agency with length of zero!")
	}
	
	for _, stoptime := range stopTimes {
		assert.NotNil(t, stoptime.TripID)
		assert.NotNil(t, stoptime.Arrival)
		assert.NotNil(t, stoptime.Departure)
		assert.NotNil(t, stoptime.StopID)
		assert.NotNil(t, stoptime.StopSequence)
	}
}


