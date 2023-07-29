package domain
/* Models based on spec: https://developers.google.com/transit/	gtfs/reference
 Currently only support primitive types, because I use golang reflect package to unmarshal

 I lump required and conditionally required together
 (until I see everythibg breaks)

 TODO:
 Change type for int and float
 Capture enum types through int -> enum definition
 See if we can doe something clever with reflect and structField tags to 
 simplify the calls throughout
*/ 

import (
	"github.com/twpayne/go-geom"	
)

type LoadProgress struct {
	Filename	string
	Percent 	int
	RowLength	int
	Index		int
	Message 	string
	Done		bool
}

type GTFS struct {
	Agencies 		[]Agency		`csv:"agency.txt" required:"true"`
	Stops 			[]Stop			`csv:"stops.txt" required:"true"`
	Routes 			[]Route 		`csv:"routes.txt" required:"true"`
	Trips 			[]Trip			`csv:"trips.txt" required:"true"`
	StopTimes 		[]StopTime		`csv:"stop_times.txt" required:"true"`
	Calendar 		[]Calendar		`csv:"calendar.txt" required:"true"`
	CalendarDates	[]CalendarDate	`csv:"calendar_dates.txt" required:"true"`
	FareAttributes 	[]FareAttribute	`csv:"fare_attributes.txt" required:"false"`
	FareRules		[]FareRule 		`csv:"fare_rules.txt" required:"false"`
	Shapes 			[]Shape 		`csv:"shapes.txt" required:"false"`
	Frequencies		[]Frequency 	`csv:"frequencies.txt" required:"false"`
	Transfers 		[]Transfer 		`csv:"transfers.txt" required:"false"`
	Pathways 		[]Pathway		`csv:"pathways.txt" required:"false"`
	Levels 			[]Level			`csv:"levels.txt" required:"false"`
	FeedInfo 		[]FeedInfo 		`csv:"feed_info.txt" required:"false"`
	Translations 	[]Translation 	`csv:"translations.txt" required:"false"`
	Attributions 	[]Attribution	`csv:"attributions.txt" required:"false"`
}

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
	GeoPoint geom.Point // I could use wbk.Point, but that is for encoding/decoding primarily. Until I know more I stick witht he more basic geom.Point
	Lat					float64 `csv:"stop_lat" required:"true"`
	Lon    				float64 `csv:"stop_lon" required:"true"`
	ZoneID				string `csv:"zone_id" required:"false"`   // only required if having fare_rules.txt in the dataset. We need some way of validating that.
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
	ID       	string 	`csv:"route_id" required:"true"`
	AgencyID	string 	`csv:"agency_id" required:"true"`
	Name		string 	`csv:"route_short_name" required:"true"`
	LongName	string 	`csv:"route_long_name" required:"true"`
	Description	string 	`csv:"route_desc" required:"false"`
	Type 		string 	`csv:"route_type" required:"true"`
	URL			string 	`csv:"route_url" required:"false"`
	Color		string 	`csv:"route_color" required:"false"`
	TextColor	string 	`csv:"route_text_color" required:"false"`
	SortOrder	int 	`csv:"route_sort_order" required:"false"`
	ContPickup	string 	`csv:"continuous_pickup" required:"false"`
	ContDrop	string 	`csv:"continuous_drop_off" required:"false"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#tripstxt
type Trip struct {
	RouteID					string `csv:"route_id" required:"true"`
	ServiceID				string `csv:"service_id" required:"true"`
	ID						string `csv:"trip_id" required:"true"`
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
	FareID				string `csv:"fare_id" required:"true"`
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

type Shape struct {
	ID					string `csv:"shape_id" required:"true"`
	// this is weird that it is required
	Lat					float64 `csv:"shape_pt_lat" required:"true"`
	Lon					float64 `csv:"shape_pt_lon" required:"true"`
	Sequence			string `csv:"shape_pt_sequence" required:"true"`
	GeoLineString		geom.LineString
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
// Spec: https://developers.google.com/transit/gtfs/reference#pathwaystxt
type Pathway struct {
	ID					string `csv:"pathway_id" required:"true"`
	FromStopID			string `csv:"from_stop_id" required:"true"`
	ToStopId			string `csv:"to_stop_id" required:"true"`
	Mode				string `csv:"pathway_mode" required:"true"`
	IsBidirectional		string `csv:"is_bidirectional" required:"true"`
	Length				string `csv:"length" required:"false"`
	TraversalTime		string `csv:"traversal_time" required:"false"`
	StairCount			string `csv:"stair_count" required:"false"`
	MaxSlope			string `csv:"max_slope" required:"false"`
	MinWidth			string `csv:"min_width" required:"false"`
	Signposted			string `csv:"signpost_as" required:"false"`
	ReversedSignposted 	string `csv:"reversed_signposted_as" required:"false"`
}

// Spec: https://developers.google.com/transit/gtfs/reference#levelstxt
type Level struct {
	ID		string `csv:"level_id" required="true"`
	Index	string `csv:"level_index" required="true"`
	Name	string `csv:"level_name" required="false"`
}
 
// Spec: https://developers.google.com/transit/gtfs/reference#feed_infotxt
type FeedInfo struct {
	PublisherName	string `csv:"feed_publisher_name" required:"true"`
	PublisherURL	string `csv:"feed_publisher_url" required:"true"`
	Language		string `csv:"feed_lang" required:"true"`
	DefaultLanguage	string `csv:"default_lang" required:"false"`
	StartDate		string `csv:"feed_start_date" required:"false"`
	EndDate			string `csv:"feed_end_date" required:"false"`
	Version			string `csv:"feed_version" required:"false"`
	ContactEmail	string `csv:"feed_contact_email" required:"false"`
	ContactURL		string `csv:"feed_contact_url" required:"false"`
}


// Spec: https://developers.google.com/transit/gtfs/reference#translationstxt
type Translation struct {
	TableName		string `csv:"table_name" required:"true"`
	FieldName		string `csv:"field_name" required:"true"`
	Language		string `csv:"language" required:"true"`
	Translation		string `csv:"translation" required:"true"`
	RecordID		string `csv:"record_id" required:"true"`
	SubRecordID		string `csv:"record_sub_id" required:"true"`
	FieldValue		string `csv:"field_value" required:"true"`			
}

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