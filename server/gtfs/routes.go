package gtfs

// Spec: https://developers.google.com/transit/gtfs/reference#routestxt
type Route struct {
	ID       	string
	AgencyID	string
	Name		string
	LongName	string
	Description	string
	Type 		Type 
	URL			string
	Color		string
	TextColor	string
	SortOrder	string
	ContPickup	ContDrop
	ContDrop	ContDrop
}

type Type int

const (
	LightRail 	Type = 0
	Metro		Type = 1
	Rail		Type = 2
	Bus			Type = 3
	Ferry		Type = 4
	CableTram	Type = 5
	ArialLift	Type = 6
	Funicular	Type = 7
	TrollyBus	Type = 11
	Monorail	Type = 12
)

type ContDrop int

const (
	Stop	ContDrop = 0
	Empty	ContDrop = 1
	Contact	ContDrop = 2
	Ask		ContDrop = 3
)

var (
	//We use the following map for checking validity and expected fields
	//true indicates required field per spec and false optional field.
	//see  https://developers.google.com/transit/gtfs/reference#agencytxt
	routeFields = map[string]bool{
		"route_id": true,
		"agency_id": true,
		"agency_url": true,
		"agency_timezone": true,
		"agency_lang": false,
		"agency_phone": false,
		"agency_fare_url": false,
		"agency_email": false,
	}
)

