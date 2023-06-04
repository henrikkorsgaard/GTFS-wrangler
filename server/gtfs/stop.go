package gtfs

import (
	//"errors"
	"strings"
	
)

// Spec: https://developers.google.com/transit/gtfs/reference#routestxt
type StopTime struct {
	TripID      		string `csv:"trip_id" required:"true"`
	Arrival				string `csv:"arrival_time" required:"true"`
	Departure			string `csv:"departure_time" required:"true"`
	StopID				string `csv:"stop_id" required:"true"`
	StopSequence		string `csv:"stop_sequence" required:"true"`
	Headsign 			string `csv:"stop_headsign" required:"false"` // Do we care --> it goes back into json anyways?
	Pickup				string `csv:"pickup_type" required:"false"`
	Dropoff				string `csv:"drop_off_type" required:"false"`
	TextColor			string `csv:"route_text_color" required:"false"`
	
	ContPickup			string `csv:"continuous_pickup" required:"false"` // Do we care --> it goes back into json anyways?
	ContDrop			string `csv:"continuous_drop_off" required:"false"` // Do we care --> it goes back into json anyways?
	DistanceTravelled	string `csv:"shape_dist_traveled" required:"false"`
	TimePoint			string `csv:"timepoint" required:"false"`
}

// Doing a dedicated instantiation will be so much faster! like factor 6
func loadStopTimes(filepath string) (stopTimes []StopTime, err error){

	rows, err := loadFromCSVFilePath(filepath)
	header := rows[0]

	// Map header index
	// We need this to create the mapping between a row value and struct field
	headerIndex := make(map[string]int)
	for i, name := range header {
		headerIndex[strings.TrimSpace(name)] = i
	}
	

	tripPos, _ := headerIndex["trip_id"]
	arrivalPos, _ := headerIndex["trip_id"]
	departPos, _ := headerIndex["trip_id"]
	stopPos, _ := headerIndex["trip_id"]
	stopSeqPos, _ := headerIndex["trip_id"]

	for _, row := range rows[1:] {
		stopTime := StopTime{
			TripID: row[tripPos],
			Arrival: row[arrivalPos],
			Departure: row[departPos],
			StopID: row[stopPos],
			StopSequence: row[stopSeqPos],
		}
		/*
		if i == 0 && !hasValidHeaderFields(rowmap, &StopTime{}) {
			err = errors.New("Agency.txt does not contain the required fields!")
			break
		// we gonna skip the first row
		} else if i == 0 {
			continue
		}*/

		/*
		stopTime := StopTime{}
		err := unmarshal(rowmap, &stopTime)
		if err != nil {
			break
		}*/
		stopTimes = append(stopTimes, stopTime)
	}
	return 
}


// Doing a dedicated instantiation will be so much faster! like factor 6
func loadStopTimesSlice(filepath string) (stopTimes []StopTime, err error){

	progress := make(chan int)
	errs := make(chan error)

	rows, err := loadFromCSVFilePath(filepath)
	header := rows[0]
	rows = rows[1:]
	//stopTimes = []StopTime{}
	go unmarshalSlice(header, rows, &stopTimes, progress, errs)

	for {
		done := false
		select {
			case err = <- errs:
				done = true
			case p := <- progress:
				if p == -1 {
					done = true
					break
				} 

				//fmt.Printf("LoadStopTimesSlice is %d percent done!\n", p)
		}
		if done {
			break
		}
	}

	return 
}