package gtfs

/* Models based on spec: https://developers.google.com/transit/	gtfs/reference
 Currently only support primitive types, because I use golang reflect package to unmarshal

 I lump required and conditionally required together
 (until I see everythibg breaks)
*/ 

//Map containing filename as string and boolean indicating if file is required or optional per the spec
var gtfsFilesRequirements = make(map[string]bool)
gtfsFilesRequirements["agency.txt"] = true
gtfsFilesRequirements["stops.txt"] = true
gtfsFilesRequirements["routes.txt"] = true
gtfsFilesRequirements["trips.txt"] = true
gtfsFilesRequirements["stop_times.txt"] = true
gtfsFilesRequirements["calendar.txt"] = true //Conditional, but will put this as required until I know if everything breaks
gtfsFilesRequirements["calendar_dates.txt"] = true //Conditional
gtfsFilesRequirements["fare_attributes.txt"] = false 
gtfsFilesRequirements["fare_rules.txt"] = false 
gtfsFilesRequirements["shapes.txt"] = false 
gtfsFilesRequirements["frequencies.txt"] = false 
gtfsFilesRequirements["transfers.txt"] = false 
gtfsFilesRequirements["pathways.txt"] = false 
gtfsFilesRequirements["levels.txt"] = false 
gtfsFilesRequirements["feed_info.txt"] = false 
gtfsFilesRequirements["translations.txt"] = false 
gtfsFilesRequirements["attributions.txt"] = false 

// Spec: https://developers.google.com/transit/gtfs/reference#attributionstxt
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

// Spec: https://developers.google.com/transit/gtfs/reference#stopstxt
type Stop struct {
	ID					string `csv:"stop_id" required:"true"`
	Code	 			string `csv:"stop_code" required:"false"`
	Name				string `csv:"stop_name" required:"true"`
	Description			string `csv:"stop_desc" required:"false"`
	Lat					string `csv:"stop_lat" required:"true"`
	Lon    				string `csv:"stop_lon" required:"true"`
	ZoneId				string `csv:"zone_id" required:"true"`
	URL 				string `csv:"stop_url" required:"false"`
	LocationType 		string `csv:"location_type" required:"false"`
	ParentStation		string `csv:"parent_station" required:"true"`
	Timezone			string `csv:"stop_timezone" required:"false"`
	WheelchairBoarding	string `csv:"wheelchair_boarding" required:"false"`
	LevelID				string `csv:"level_id" required:"false"`
	PlatformCode		string `csv:"platform_code" required:"false"`
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

// Spec: https://developers.google.com/transit/gtfs/reference#tripstxt
type Trip struct {
	RouteID					string `csv:"route_id" required:"true"`
	ServiceID				string `csv:"service_id" required:"true"`
	TripID					string `csv:"trip_id" required:"true"`
	TripHeadsign			string `csv:"trip_headsign" required:"false"`
	Name					string `csv:"trip_short_name" required:"false"`
	DirectionID 			string `csv:"direction_id" required:"false"`
	BlockID					string `csv:"block_id" required:"false"`
	ShapeID					string `csv:"shape_id" required:"true"`
	WheelchairAccessible	string `csv:"wheelchair_accessible" required:"false"`
	BikesAllowed			string `csv:"bikes_allowed" required:"false"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#routestxt
// Note this does not have an ID, as it is a relational type (TripID - StopID)
type StopTime struct {
	TripID      		string `csv:"trip_id" required:"true"`
	Arrival				string `csv:"arrival_time" required:"true"`
	Departure			string `csv:"departure_time" required:"true"`
	StopID				string `csv:"stop_id" required:"true"`
	StopSequence		string `csv:"stop_sequence" required:"true"`
	StopHeadsign 		string `csv:"stop_headsign" required:"false"`
	Pickup				string `csv:"pickup_type" required:"false"`
	Dropoff				string `csv:"drop_off_type" required:"false"`
	TextColor			string `csv:"route_text_color" required:"false"`
	ContPickup			string `csv:"continuous_pickup" required:"false"`
	ContDrop			string `csv:"continuous_drop_off" required:"false"`
	DistanceTravelled	string `csv:"shape_dist_traveled" required:"false"`
	TimePoint			string `csv:"timepoint" required:"false"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#calendartxt

type Calendar struct {
	ServiceID	string `csv:"service_id" required:"true"`
	Monday		string `csv:"monday" required:"true"`
	Tuesday		string `csv:"tuesday" required:"true"`
	Wednesday	string `csv:"wednesday" required:"true"`
	Thursday	string `csv:"thursday" required:"true"`
	Friday		string `csv:"friday" required:"true"`
	Saturday	string `csv:"saturday" required:"true"`
	Sunday		string `csv:"sunday" required:"true"`
	StartDate	string `csv:"start_date" required:"true"`
	EndDate 	string `csv:"end_date" required:"true"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#calendar_datestxt

type CalendarDate struct {
	ServiceID 	string `csv:"service_id" required:"true"`
	Date		string `csv:"date" required:"true"`
	Exception	string `csv:"exception_type" required:"true"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#fare_attributestxt

type FareAttribute struct {
	FareID					string `csv:"fare_id" required:"true"`
	Price				string `csv:"price" required:"true"`
	Currency 			string `csv:"currency_type" required:"true"`
	PaymentMethod		string `csv:"payment_method" requied:"true"`
	Transfers			string `csv:"transfers" required:"true"`
	AgencyID			string `csv:"agency_id" required:"true"`
	TransferDuration	string `csv:"transfer_duration" required:"true"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#fare_rulestxt

type FareRule struct {
	FareID			string `csv:"fare_id" required:"true"`
	RouteID			string `csv:"route_id" required:"false"`
	OriginId		string `csv:"origin_id" required:"false"`
	DestinationID	string `csv:"destination_id" required:"false"`
	ContainsId		string `csv:"contains_id" required:"false"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#shapestxt

type Shape {
	ID					string `csv:"shape_id" required:"true"`
	Lat					string `csv:"shape_pt_lat" required:"true"`
	Lon					string `csv:"shape_pt_lon" required:"true"`
	Sequence			string `csv:"shape_pt_sequence" required:"true"`
	DistanceTraveled	string `csv:"shape_dist_traveled" required:"false"`
}


// Spec: https://developers.google.com/transit/gtfs/reference#frequenciestxt

type Frequency struct {
	TripID		string `csv:"trip_id" required="true"`
	StartTime	string `csv:"start_time" required="true"`
	EndTime		string `csv:"end_time" required="true"`
	HeadwaySec	string `csv:"headway_secs" required="true"`
	ExactTimes	string `csv:"exact_times" required="true"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#transferstxt

type Transfer struct {
	FromStopID		string `csv:"from_stop_id" required="true"`
	ToStopID		string `csv:"to_stop_id" required="true"`
	Type			string `csv:"transfer_type" required="true"`
	MinTransferTime	string `csv:"min_transfer_time" required="false"`
}

//TODO PATHWAYS 
// Spec: https://developers.google.com/transit/gtfs/reference#pathwaystxt

//TODO LEVELS 
// Spec: https://developers.google.com/transit/gtfs/reference#levelstxt

//TODO FEED_INFO 
// Spec: https://developers.google.com/transit/gtfs/reference#feed_infotxt

//TODO TRANSLATIONS
// Spec: https://developers.google.com/transit/gtfs/reference#translationstxt


// Spec: https://developers.google.com/transit/gtfs/reference#agencytxt
type Attribution struct {
	ID				string `csv:"attribution_id" required:"false"`
	AgencyID		string `csv:"agency_id" required:"false"`
	RouteID	 		string `csv:"route_id" required:"false"`
	TripID			string `csv:"trip_id" required:"false"`
	Organization	string `csv:"organization_name" required:"true"`
	IsProducer		string `csv:"is_producer" required:"false"`
	IsOperator    	string `csv:"is_operator" required:"false"`
	IsAuthority    	string `csv:"is_authority" required:"false"`
	URL				string `csv:"attribution_url" required:"false"`
	Email			string `csv:"attribution_email" required:"false"`
	Phone			string `csv:"attribution_phone" required:"false"`
}