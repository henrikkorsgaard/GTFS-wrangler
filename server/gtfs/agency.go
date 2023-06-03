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


// we could change this to take the data, and then an "any" paramter as a reference. Then return it when done?
// TODO tomorrow
func unmarshal(data map[string]string, obj any) (err error) {

	vof := reflect.ValueOf(obj)
	tof := reflect.TypeOf(obj)
	
	//we likely miss some checks here
	//we want to ensure that kind is pointer (ptr) to struct
	if vof.Kind() != reflect.Pointer || vof.IsNil() {
		err = errors.New("Cannot unmarshal " + tof.Name() + ". Needs to be a pointer to a struct!")
		return
	}
	
	fields := reflect.VisibleFields(tof.Elem())
	for _, field := range fields {
		
		csvKey := field.Tag.Get("csv")
		required := field.Tag.Get("required")
		isRequired, err := strconv.ParseBool(required)
		if err != nil {
			break
		}

		_, ok :=  data[csvKey]

		if isRequired && !ok {
			// Need to refine this to be more generalised
			err = errors.New("Unable to unmarshal data into " + tof.Name() + ". Missing required fields according to GTFS specification.")
			break
		}
		
		vof.Elem().FieldByName(field.Name).SetString(data[csvKey])
	}
	
	return
}

//This need to be generalized as well
func loadAgenciesFromCSVFilePath(filepath string) (agencies []Agency, err error){

	csvfile, err := os.Open(filepath)
	if err != nil {
		return
	}

	r := csv.NewReader(csvfile)
	all, err := r.ReadAll()

	if err != nil {
		// we need a lib specific error
		return
	}

	//if we want to return progres
	//we need a chan and then return per row read
	for i, row := range all {
		rowmap := map[string]string{}
		for i, item := range row {
			rowmap[all[0][i]] = item
		}

		if i == 0 && !isValidHeaderFields(rowmap) {
			err = errors.New("Agency.txt does not contain the required fields!")
			break
		}
		agency := Agency{}
		err := unmarshal(rowmap, &agency)
		if err != nil {
			break
		}
		agencies = append(agencies, agency)
	}

	return
}


// need to generalize as well
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

