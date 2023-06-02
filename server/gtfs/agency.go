package gtfs

import (
	"os"
	"bufio"
	"strings"
	"errors"	
)

var (
	//We use the following map for checking validity and expected fields
	//true indicates required field per spec and false optional field.
	//see  https://developers.google.com/transit/gtfs/reference#agencytxt
	agencyFields = map[string]bool{
		"agency_id": true,
		"agency_name": true,
		"agency_url": true,
		"agency_timezone": true,
		"agency_lang": false,
		"agency_phone": false,
		"agency_fare_url": false,
		"agency_email": false,
	}
)

type Agency struct {
	ID       	string
	Name     	string
	URL			string
	Timezone	string
	Lang		string
	Phone    	string
}

// I don't want to use some of the csv unmarshall libraries because they seem to be buggy
// I know the file and the target struct, so this should just be a dedicated unmarshal
func loadAgenciesFromFilePath(filepath string) (agencies []Agency, err error) {

	lines, err := readFileIntoLines(filepath)
	if err != nil {
		return
	}

	stripped := strings.ReplaceAll(lines[0], "\"", "")
	fields := strings.Split(stripped, ",")
	
	if(!hasRequiredAgencyFields(fields)){
		err = errors.New("Invalid agency.txt file. Headers does not match expectation!")
		return
	}

	fieldIndexes := getValidFieldColumnIndex(fields)

	// we skip first line because it's headers!
	for _, row := range lines[1:]{
		stripped = strings.ReplaceAll(row, "\"", "")
		vals := strings.Split(stripped, ",")

		// This is somewhat ceremonial to flag for data inconcistency in header/row
		if len(vals) != len(fieldIndexes) {
			err = errors.New("Error: Mismatch between number of fields and row index!")
			return
		}
		
		agency := Agency{
			ID: vals[fieldIndexes["agency_id"]],
			Name: vals[fieldIndexes["agency_name"]],
			URL: vals[fieldIndexes["agency_url"]],
			Timezone: vals[fieldIndexes["agency_timezone"]],
		}

		if _, ok := fieldIndexes["agency_lang"]; ok {
			agency.Lang = vals[fieldIndexes["agency_lang"]]
		}

		if _, ok := fieldIndexes["agency_phone"]; ok {
			agency.Lang = vals[fieldIndexes["agency_phone"]]
		}

		if _, ok := fieldIndexes["agency_fare_url"]; ok {
			agency.Lang = vals[fieldIndexes["agency_fare_url"]]
		}

		if _, ok := fieldIndexes["agency_email"]; ok {
			agency.Lang = vals[fieldIndexes["agency_email"]]
		}

		agencies = append(agencies, agency)

	}
	
	return
}

func readFileIntoLines(filepath string) (lines []string, err error) {

	readFile, err := os.Open(filepath)
    if err != nil {
        return
    }
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    
  
    for fileScanner.Scan() {
        lines = append(lines, fileScanner.Text())
    }
  
    return
}

// We cannot assume that GTFS files are ordered according to spec.
func getValidFieldColumnIndex(fields []string) (fieldIndexes map[string]int) {

	fieldIndexes = make(map[string]int)

	for index, field := range fields {
		//we do not care about fields that are not in agencyFields
		//this is a side effect call essentially ignoring invalid fields!
		if _, ok := agencyFields[field]; ok {
			fieldIndexes[field] = index
		}
	}

	return
}

// currently just checking that the fields in the file are among the fields in the agency spec.
// We need to check for required fields first, then for valid fields next
func hasRequiredAgencyFields(fields []string) (valid bool) { 

	// we do not cover the required vs the optional vs unexpected fields.
	requiredFields := make(map[string]bool)
	requiredCount := 0
	for k, v := range agencyFields {
		if v {
			requiredFields[k] = false
		}
	}

	for _, field := range fields {
		if _, ok := requiredFields[field];ok {
			requiredCount++
		}
	}
	
	if len(requiredFields) == requiredCount {
		valid = true
		
	}
	
	return
}


//util function that is useful for more than one file
func contains(stringSlice []string, target string) bool {
	for _, s := range stringSlice {
		if s == target {
			return true
		}
	}
	return false
}
