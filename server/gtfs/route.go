package gtfs

import (
	"errors"
)

// Spec: https://developers.google.com/transit/gtfs/reference#routestxt
type Route struct {
	ID       	string `csv:"route_id" required:"true"`
	AgencyID	string `csv:"agency_id" required:"true"`
	Name		string `csv:"route_short_name" required:"true"`
	LongName	string `csv:"route_long_name" required:"true"`
	Description	string `csv:"route_desc" required:"false"`
	Type 		RouteType `csv:"route_type" required:"true"` // Do we care --> it goes back into json anyways?
	URL			string `csv:"route_url" required:"false"`
	Color		string `csv:"route_color" required:"false"`
	TextColor	string `csv:"route_text_color" required:"false"`
	SortOrder	string `csv:"route_sort_order" required:"false"`
	ContPickup	string `csv:"continuous_pickup" required:"false"` // Do we care --> it goes back into json anyways?
	ContDrop	string `csv:"continuous_drop_off" required:"false"` // Do we care --> it goes back into json anyways?
}

type RouteType int

const (
	LightRail 	RouteType = 0
	Metro		RouteType = 1
	Rail		RouteType = 2
	Bus			RouteType = 3
	Ferry		RouteType = 4
	CableTram	RouteType = 5
	ArialLift	RouteType = 6
	Funicular	RouteType = 7
	TrollyBus	RouteType = 11
	Monorail	RouteType = 12
)

//yes, I could use stringer, but this is fucking artisinal software or convivial computing or whatever
func (rt RouteType) String() string {

	return []string{"LightRail", "Metro", "Rail", "Bus", "Ferry", "CableTram","ArialLift","Funicular","unknown route type","unknown route type","unknown route type","TrollyBus", "Monorail"}[rt]
}

func loadRoutes(filepath string) (routes []Route, err error){
	rows, err := loadFromCSVFilePath(filepath)

	//if we want to return progres
	//we need a chan and then return per row read
	for i, row := range rows {
		rowmap := map[string]string{}
		for i, item := range row {
			rowmap[rows[0][i]] = item
		}

		if i == 0 && !hasValidHeaderFields(rowmap, &Route{}) {
			err = errors.New("Agency.txt does not contain the required fields!")
			break
		// we gonna skip the first row
		} else if i == 0 {
			continue
		}
		route := Route{}
		err := unmarshal(rowmap, &route)
		if err != nil {
			break
		}
		routes = append(routes, route)
	}
	return 
}