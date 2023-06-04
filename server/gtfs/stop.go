package gtfs

import (

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
func loadStopTimesSlice(filepath string) (stopTimes []StopTime, err error){

	progress := make(chan int)
	errs := make(chan error)

	rows, err := loadFromCSVFilePath(filepath)

	if err != nil {
		
		return
	}
	
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