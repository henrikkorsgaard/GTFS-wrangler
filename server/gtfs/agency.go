package gtfs

import (
	//"errors"
	
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
/*
func loadAgencies(filepath string) (agencies []Agency, err error){
	rows, err := loadFromCSVFilePath(filepath)

	//if we want to return progres
	//we need a chan and then return per row read
	for i, row := range rows {
		rowmap := map[string]string{}
		for i, item := range row {
			rowmap[rows[0][i]] = item
		}

		if i == 0 && !hasValidHeaderFields(rowmap, &Agency{}) {
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
}*/




