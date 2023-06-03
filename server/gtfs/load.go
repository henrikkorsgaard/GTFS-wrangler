package gtfs

import (
	"reflect"
	"errors"
	"os"
	"strconv"
	"encoding/csv"

	"fmt"
)

/*
	Unmarshal function tailored to used with the CSV package
	and the GTFS structs. Does not account for edge-cases and whatnot
*/
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

		/*
			We will have to compare with .Type() because Kind() gives the underlying type and we want to check up against multiple int enums. Using the .Kind() and reflect.String comparions will not allow us to check for enums specifically.

			This is also the techniques (numtype) in the encoding/json unmarshal lib, see:
			https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/encoding/json/decode.go;drc=46ab7a5c4f80d912f25b6b3e1044282a2a79df8b;l=964


			https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/encoding/json/decode.go;drc=46ab7a5c4f80d912f25b6b3e1044282a2a79df8b;l=855

			Alternative strategy: We could check for ints first and then catch all the individual cases. But that would introduce one additional layer of abstraction to serve some ideomatic cleaness on .Kind() == reflect.String checks!

		*/	

		var routetype = reflect.TypeOf(RouteType(0))
		var stringtype = reflect.TypeOf("")

		objField := vof.Elem().FieldByName(field.Name)

		switch objField.Type() {
			case routetype:
				rtint, err := strconv.Atoi(data[csvKey])
				if err != nil {
					break
				}
				obj.(*Route).Type = RouteType(rtint) //DIRTY BUT FIXES IT!
			case stringtype:
				objField.SetString(data[csvKey])
			default:
				fmt.Println(objField.Type())
				fmt.Println(data[csvKey])
				fmt.Println("never, NEVER !")
		}
	}
	
	return
}

//This need to be generalized as well
func loadFromCSVFilePath(filepath string) (data [][]string, err error){

	csvfile, err := os.Open(filepath)
	if err != nil {
		return
	} // DO WE CLOSE THIS?

	r := csv.NewReader(csvfile)
	data, err = r.ReadAll()

	return
}

// We run this loop for each field. That is 4.000.000 X fields checks
func hasValidHeaderFields(header map[string]string, obj any)(valid bool){
	valid = true
	tof := reflect.TypeOf(obj)
	fields := reflect.VisibleFields(tof.Elem())
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