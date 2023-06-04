package gtfs

// Spec: https://developers.google.com/transit/gtfs/reference#agencytxt
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

// Spec: https://developers.google.com/transit/gtfs/reference#routestxt
type Route struct {
	ID       	string `csv:"route_id" required:"true"`
	AgencyID	string `csv:"agency_id" required:"true"`
	Name		string `csv:"route_short_name" required:"true"`
	LongName	string `csv:"route_long_name" required:"true"`
	Description	string `csv:"route_desc" required:"false"`
	Type 		string `csv:"route_type" required:"true"`
	URL			string `csv:"route_url" required:"false"`
	Color		string `csv:"route_color" required:"false"`
	TextColor	string `csv:"route_text_color" required:"false"`
	SortOrder	string `csv:"route_sort_order" required:"false"`
	ContPickup	string `csv:"continuous_pickup" required:"false"`
	ContDrop	string `csv:"continuous_drop_off" required:"false"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#routestxt
type StopTime struct {
	TripID      		string `csv:"trip_id" required:"true"`
	Arrival				string `csv:"arrival_time" required:"true"`
	Departure			string `csv:"departure_time" required:"true"`
	StopID				string `csv:"stop_id" required:"true"`
	StopSequence		string `csv:"stop_sequence" required:"true"`
	Headsign 			string `csv:"stop_headsign" required:"false"`
	Pickup				string `csv:"pickup_type" required:"false"`
	Dropoff				string `csv:"drop_off_type" required:"false"`
	TextColor			string `csv:"route_text_color" required:"false"`
	ContPickup			string `csv:"continuous_pickup" required:"false"`
	ContDrop			string `csv:"continuous_drop_off" required:"false"`
	DistanceTravelled	string `csv:"shape_dist_traveled" required:"false"`
	TimePoint			string `csv:"timepoint" required:"false"`
}