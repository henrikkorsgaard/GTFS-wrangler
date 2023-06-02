package gtfs

import (
	"testing"
	"github.com/stretchr/testify/assert"
	
)

func TestLoadingAgencyFromFileHeaderError(t *testing.T){
	_, err := loadAgenciesFromFilePath("./test_data/agency_invalid_header.txt")
	assert.Error(t, err)
}

func TestLoadingAgencyFromFileRowError(t *testing.T){
	_, err := loadAgenciesFromFilePath("./test_data/agency_invalid_row.txt")
	assert.Error(t, err)
}


func TestLoadingAgencyFromFile(t *testing.T){


	//test data ?
	// or do we just test that each returns with information?

	agencies, err := loadAgenciesFromFilePath("./test_data/agency.txt")
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