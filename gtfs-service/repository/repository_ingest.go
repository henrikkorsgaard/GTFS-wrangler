package repository

import (
	"henrikkorsgaard.dk/gtfs-service/domain"
	"github.com/lib/pq"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
)

func (repo *repository) IngestAgency(agency []domain.Agency) (err error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("agency","id", "name", "url", "timezone", "lang", "phone", "fare_url", "email"))
	if err != nil {
		return
	}

	for _, a := range agency {	
		
		_, err = stmt.Exec(&a.ID, &a.Name,&a.URL, &a.Timezone,&a.Lang, &a.Phone, &a.FareURL,&a.Email)

		if err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	return tx.Commit()
}

func (repo *repository) IngestStops(stops []domain.Stop) (err error){
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("stops","id","code", "name","description","location","zone_id","url", "location_type","parent_station","timezone","wheelchair_boarding", "level_id", "platform_code"))
	if err != nil {
		return
	}

	for _, s := range stops {	
		ewkbhexGeom, err := ewkbhex.Encode(&s.GeoPoint, ewkbhex.NDR)
		if err != nil {
			return err
		}

		if _, err = stmt.Exec(&s.ID, &s.Code, &s.Name, &s.Description,ewkbhexGeom, &s.ZoneID,&s.URL, &s.LocationType, &s.ParentStation, &s.Timezone, &s.WheelchairBoarding,&s.LevelID, &s.PlatformCode); err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	return tx.Commit()
}

func (repo *repository) IngestRoutes(routes []domain.Route) (err error){
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("routes","id", "agency_id", "short_name", "long_name", "description","type", "url","color","text_color","sort_order","continuous_pickup", "continuous_drop_off"))
	if err != nil {
		return
	}

	for _, r := range routes {	
		
		if _, err = stmt.Exec(r.ID, r.AgencyID, r.Name, r.LongName, r.Description, r.Type, r.URL, r.Color, r.TextColor, r.SortOrder, r.ContPickup, r.ContDrop); err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	return tx.Commit()
}

func (repo *repository) IngestTrips(trips []domain.Trip) (err error){
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("trips","id","route_id", "service_id", "shape_id", "headsign", "name","wheelchair_accessible", "bikes_allowed"))
	if err != nil {
		return
	}

	for _, t := range trips {	
		
		if _, err = stmt.Exec(t.ID,t.RouteID,t.ServiceID, t.ShapeID, t.Headsign, t.Name, t.WheelchairAccessible,t.BikesAllowed); err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	return tx.Commit()
}

func (repo *repository) IngestShapes(shapes []domain.Shape) (err error){
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("shapes","id", "geo_line"))
	if err != nil {
		return
	}

	for _, s := range shapes {	
		ewkbhexGeom, err := ewkbhex.Encode(&s.GeoLineString, ewkbhex.NDR)
		if err != nil {
			return err
		}

		if _, err = stmt.Exec(s.ID, ewkbhexGeom); err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	return tx.Commit()
}

func (repo *repository) IngestStopTimes(stopTimes []domain.StopTime) (err error){
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("stoptimes","trip_id", "stop_id", "arrival", "departure", "stop_sequence", "stop_headsign", "pickup_type","drop_off_type","continuous_pickup", "continuous_drop_off", "shape_dist_traveled", "timepoint"))
	if err != nil {
		return
	}

	for _, st := range stopTimes {	
		
		if _, err = stmt.Exec(st.TripID, st.StopID, st.Arrival, st.Departure, st.StopSequence, st.StopHeadsign, st.Pickup, st.Dropoff,st.ContPickup,st.ContDrop, st.DistanceTraveled, st.Timepoint); err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	return tx.Commit()
}

func (repo *repository) IngestCalendars(calendars []domain.Calendar) (err error){
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("calendar","service_id", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday","sunday","start_date", "end_date"))
	if err != nil {
		return
	}

	for _, c := range calendars {	
		
		if _, err = stmt.Exec(c.ServiceID, c.Monday, c.Tuesday, c.Wednesday, c.Thursday, c.Friday, c.Saturday, c.Sunday, c.StartDate, c.EndDate); err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	return tx.Commit()
}

