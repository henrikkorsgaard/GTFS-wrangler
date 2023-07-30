package repository

import (
	"fmt"
	"henrikkorsgaard.dk/gtfs-service/domain"
	"github.com/twpayne/go-geom/encoding/ewkb"
)


func (repo *repository) FetchAgency() (agency []domain.Agency, err error){
	rows, err := repo.db.Query("SELECT id, name, url, timezone, lang, phone, fare_url, email FROM agency;")
	defer rows.Close()

	if err != nil {
		return
	}
	
	for rows.Next() {
		a := domain.Agency{}
	
		err = rows.Scan(&a.ID, &a.Name, &a.URL, &a.Timezone, &a.Lang, &a.Phone, &a.FareURL, &a.Email)
		if err != nil {
			break
		}
	
		agency = append(agency, a)
	}

	if rows.Err() != nil {
		return 
	}

	return
}

func (repo *repository) FetchStops() (stops []domain.Stop, err error){
	rows, err := repo.db.Query("SELECT id, code, name, description, ST_AsBinary(location), zone_id, url, location_type, parent_station, timezone, wheelchair_boarding, level_id, platform_code FROM stops;")
	defer rows.Close()

	if err != nil {
		return
	}
	
	for rows.Next() {
		s := domain.Stop{}
		var p ewkb.Point
		err = rows.Scan(&s.ID, &s.Code, &s.Name, &s.Description,&p, &s.ZoneID,&s.URL, &s.LocationType, &s.ParentStation, &s.Timezone, &s.WheelchairBoarding,&s.LevelID, &s.PlatformCode)
		if err != nil {
			break
		}
		s.GeoPoint = *p.Point
		stops = append(stops, s)
	}

	if rows.Err() != nil {
		return 
	}

	return
}

func (repo *repository) FetchRoutes() (routes []domain.Route, err error){
	rows, err := repo.db.Query("SELECT id, agency_id, short_name, long_name, description,  type, url, color, text_color, sort_order, continuous_pickup, continuous_drop_off FROM routes;")
	defer rows.Close()

	if err != nil {
		return
	}
	
	for rows.Next() {
		r := domain.Route{}
		
		err = rows.Scan(&r.ID, &r.AgencyID, &r.Name, &r.LongName, &r.Description,&r.Type,&r.URL, &r.Color, &r.TextColor, &r.SortOrder, &r.ContPickup, &r.ContDrop)
		if err != nil {
			break
		}
		
		routes = append(routes, r)
	}

	if rows.Err() != nil {
		return 
	}

	return
}

func (repo *repository) FetchTrips() (trips []domain.Trip, err error){
	rows, err := repo.db.Query("SELECT id, route_id, service_id, shape_id, headsign, name, block_id, wheelchair_accessible, bikes_allowed FROM trips;")
	defer rows.Close()
	
	if err != nil {
		return
	}

	for rows.Next() {
		t := domain.Trip{}
		
		err = rows.Scan(&t.ID, &t.RouteID,&t.ServiceID,&t.ShapeID, &t.Headsign, &t.Name, &t.BlockID, &t.WheelchairAccessible, &t.BikesAllowed)
		if err != nil {
			fmt.Println(err)
			break
		}
		
		trips = append(trips, t)
	}

	if rows.Err() != nil {
		return 
	}

	return
}

func (repo *repository) FetchShapes() (shapes []domain.Shape, err error){
	rows, err := repo.db.Query("SELECT id, ST_AsBinary(geo_line) FROM shapes;")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		
		s := domain.Shape{}
		var ls ewkb.LineString
		err = rows.Scan(&s.ID, &ls)
		if err != nil {
			break
		}

		s.GeoLineString = *ls.LineString
		shapes = append(shapes, s)
	}

	if rows.Err() != nil {
		return 
	}

	return
}

func (repo *repository) FetchStopTimes() (stopTimes []domain.StopTime, err error){
	rows, err := repo.db.Query("SELECT trip_id, stop_id, arrival, departure FROM stoptimes;")
	defer rows.Close()
	
	if err != nil {
		return
	}

	for rows.Next() {
		st := domain.StopTime{}
		
		err = rows.Scan(&st.TripID, &st.StopID, &st.Arrival, &st.Departure)
		if err != nil {
			fmt.Println(err)
			break
		}
		
		stopTimes = append(stopTimes, st)
	}

	if rows.Err() != nil {
		return 
	}

	return
}

