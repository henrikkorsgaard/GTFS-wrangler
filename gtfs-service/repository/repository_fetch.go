package repository

import (
	"database/sql"
	"henrikkorsgaard.dk/gtfs-service/domain"
	"github.com/twpayne/go-geom/encoding/ewkb"
)

func (repo *repository) FetchAgency() (agency []domain.Agency, err error){
	query := "SELECT id, name, url, timezone, lang, phone, fare_url, email FROM agency;"
	
	rowHandler := func(rs *sql.Rows) (err error){
		
		a := domain.Agency{}
	
		err = rs.Scan(&a.ID, &a.Name, &a.URL, &a.Timezone, &a.Lang, &a.Phone, &a.FareURL, &a.Email)
		if err != nil {
			return
		}
	
		agency = append(agency, a)
		return
	}

	err = repo.fetch(query, rowHandler)

	return
}

func (repo *repository) FetchStops() (stops []domain.Stop, err error){
	query := "SELECT id, code, name, description, ST_AsBinary(location), zone_id, url, location_type, parent_station, timezone, wheelchair_boarding, level_id, platform_code FROM stops;"
	
	rowHandler := func(rs *sql.Rows) (err error){
		
		s := domain.Stop{}
		var p ewkb.Point
		
		err = rs.Scan(&s.ID, &s.Code, &s.Name, &s.Description,&p, &s.ZoneID,&s.URL, &s.LocationType, &s.ParentStation, &s.Timezone, &s.WheelchairBoarding,&s.LevelID, &s.PlatformCode)
		if err != nil {
			return 
		}

		s.GeoPoint = *p.Point
		stops = append(stops, s)
		return
	}

	err = repo.fetch(query, rowHandler)

	return
}

func (repo *repository) FetchRoutes() (routes []domain.Route, err error){
	query := "SELECT id, agency_id, short_name, long_name, description,  type, url, color, text_color, sort_order, continuous_pickup, continuous_drop_off FROM routes;"

	rowHandler := func(rs *sql.Rows) (err error){
		r := domain.Route{}
		
		err = rs.Scan(&r.ID, &r.AgencyID, &r.Name, &r.LongName, &r.Description,&r.Type,&r.URL, &r.Color, &r.TextColor, &r.SortOrder, &r.ContPickup, &r.ContDrop)
		if err != nil {
			return
		}
		
		routes = append(routes, r)

		return
	}

	err = repo.fetch(query, rowHandler)

	return
}

func (repo *repository) FetchTrips() (trips []domain.Trip, err error){
	query := "SELECT id, route_id, service_id, shape_id, headsign, name, block_id, wheelchair_accessible, bikes_allowed FROM trips;"

	rowHandler := func(rs *sql.Rows) (err error){
		t := domain.Trip{}
		err = rs.Scan(&t.ID, &t.RouteID,&t.ServiceID,&t.ShapeID, &t.Headsign, &t.Name, &t.BlockID, &t.WheelchairAccessible, &t.BikesAllowed)
		
		if err != nil {
			return
		}

		trips = append(trips, t)

		return
	}

	err = repo.fetch(query, rowHandler)

	return
}

func (repo *repository) FetchShapes() (shapes []domain.Shape, err error){

	query := "SELECT id, ST_AsBinary(geo_line) FROM shapes;"

	rowHandler := func(rs *sql.Rows) (err error){
		s := domain.Shape{}
		var ls ewkb.LineString
		err = rs.Scan(&s.ID, &ls)
		
		if err != nil {
			return
		}

		s.GeoLineString = *ls.LineString
		shapes = append(shapes, s)
		
		return
	}

	err = repo.fetch(query, rowHandler)

	return
}

func (repo *repository) FetchStopTimes() (stopTimes []domain.StopTime, err error){
	query := "SELECT trip_id, stop_id, arrival, departure, stop_sequence, stop_headsign, pickup_type, drop_off_type, continuous_pickup, continuous_drop_off, shape_dist_traveled, timepoint FROM stoptimes;"

	rowHandler := func(rs *sql.Rows) (err error){
		st := domain.StopTime{}
		
		err = rs.Scan(&st.TripID, &st.StopID, &st.Arrival, &st.Departure, &st.StopSequence, &st.StopHeadsign, &st.Pickup, &st.Dropoff, &st.ContPickup, &st.ContDrop, &st.DistanceTraveled, &st.Timepoint)
		
		if err != nil {
			return
		}

		stopTimes = append(stopTimes, st)
		return
	}
	
	err = repo.fetch(query, rowHandler)

	return
}

func (repo *repository) FetchCalendars() (calendars []domain.Calendar, err error){
	query := "SELECT service_id, monday, tuesday, wednesday, thursday, friday, saturday, sunday, start_date, end_date FROM calendar;"

	rowHandler := func(rs *sql.Rows) (err error){
		c := domain.Calendar{}
		
		err = rs.Scan(&c.ServiceID, &c.Monday, &c.Tuesday, &c.Wednesday, &c.Thursday, &c.Friday, &c.Saturday, &c.Sunday, &c.StartDate, &c.EndDate)
		
		if err != nil {
			return
		}

		calendars = append(calendars, c)
		return
	}
	
	err = repo.fetch(query, rowHandler)

	return
}

func (repo *repository) fetch(query string, rowHandler func(r *sql.Rows) (err error)) (err error) {
	rows, err := repo.db.Query(query)
	defer rows.Close()
	
	if err != nil {
		return
	}

	for rows.Next() {
		err = rowHandler(rows)
		if err != nil {
			break
		}
	}

	if rows.Err() != nil {
		return 
	}

	return
}