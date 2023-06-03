package gtfs

import (
	//"errors"
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

	// I need to write a faster way to do this
	rows, err := loadFromCSVFilePath(filepath)

	//is it faster to do it by row?


	//if we want to return progres
	//we need a chan and then return per row read

	
	for _, row := range rows[1:] {
		rowmap := map[string]string{}
		// this doubles the runtime right there!
		for i, item := range row {
			rowmap[rows[0][i]] = item
		}

		stopTime := StopTime{
			TripID: rowmap["trip_id"],
			Arrival: rowmap["arrival_time"],
			Departure: rowmap["departure_time"],
			StopID: rowmap["stop_id"],
			StopSequence: rowmap["stop_sequence"],
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
