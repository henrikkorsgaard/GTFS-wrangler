package gtfs

import (
	"encoding/csv"
	"os"
	"errors"
	"reflect"
	"strconv"
)


type Agency struct {
	ID			string `csv:"agency_id" required:"true"`
	Name	 	string `csv:"agency_name" required:"true"`
	URL			string `csv:"agency_url" required:"true"`
	Timezone	string `csv:"agency_timezone" required:"true"`
	Lang		string `csv:"agency_lang" required:"false"`
	Phone    	string `csv:"agency_phone" required:"false"`
	FareURL		string `csv:"agency_fare_url" required:"false"`
	Email 		string `csv:"agency_email" required:"false"`
}

func unmarshalAgency(row map[string]string) (agency Agency, err error) {
	
	ref := reflect.ValueOf(&agency).Elem()
	
	fields := reflect.VisibleFields(reflect.TypeOf(agency))
	for _, field := range fields {
		
		csvKey := field.Tag.Get("csv")
		required := field.Tag.Get("required")
		isRequired, err := strconv.ParseBool(required)
		if err != nil {
			break
		}

		_, ok :=  row[csvKey]

		if isRequired && !ok {
			err = errors.New("Agency.txt does not contain rows with data for the required fields!")
			break
		}
		ref.FieldByName(field.Name).SetString(row[csvKey])
	}
	
	return
}

func loadAgenciesFromCSVFilePath(filepath string) (agencies []Agency, err error){

	csvfile, err := os.Open(filepath)
	if err != nil {
		return
	}

	r := csv.NewReader(csvfile)
	all, err := r.ReadAll()

	if err != nil {
		return
	}

	for i, row := range all {
		rowmap := map[string]string{}
		for i, item := range row {
			rowmap[all[0][i]] = item
		}

		if i == 0 && !isValidHeaderFields(rowmap) {
			err = errors.New("Agency.txt does not contain the required fields!")
			break
		}

		agency, err := unmarshalAgency(rowmap)
		if err != nil {
			break
		}
		agencies = append(agencies, agency)
	}

	return
}

func isValidHeaderFields(header map[string]string)(valid bool){
	a := Agency{}
	valid = true
	fields := reflect.VisibleFields(reflect.TypeOf(a))
	for _, field := range fields {
		csvKey := field.Tag.Get("csv")
		required := field.Tag.Get("required")
		isRequired, err := strconv.ParseBool(required)
		if err != nil {
			break
		}

		_, ok :=  header[csvKey]

		if isRequired && !ok {
			valid = false
			break
		}
	}
	
	return valid
}

