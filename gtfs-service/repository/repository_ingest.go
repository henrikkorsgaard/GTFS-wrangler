package repository

import (
	/*
	"context"
	
	"github.com/twpayne/go-geom/encoding/wkb"
	*/
	
	"henrikkorsgaard.dk/gtfs-service/domain"
	"github.com/lib/pq"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
)

func (repo *repository) IngestStops(stops []domain.Stop) (err error){
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("stops","id", "name", "description", "geo_point", "parent_station"))
	if err != nil {
		return
	}

	for _, s := range stops {	
		ewkbhexGeom, err := ewkbhex.Encode(&s.GeoPoint, ewkbhex.NDR)
		if err != nil {
			return err
		}

		if _, err = stmt.Exec(s.ID, s.Name, s.Description, ewkbhexGeom, s.ParentStation); err != nil {
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

	stmt, err := tx.Prepare(pq.CopyIn("routes","id", "agency_id", "short_name", "long_name", "type"))
	if err != nil {
		return
	}

	for _, r := range routes {	
		
		if _, err = stmt.Exec(r.ID, r.AgencyID, r.Name, r.LongName, r.Type); err != nil {
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

	stmt, err := tx.Prepare(pq.CopyIn("trips","id", "service_id", "route_id", "shape_id", "trip_headsign"))
	if err != nil {
		return
	}

	for _, t := range trips {	
		
		if _, err = stmt.Exec(t.ID,t.ServiceID,t.RouteID,t.ShapeID, t.TripHeadsign); err != nil {
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

	stmt, err := tx.Prepare(pq.CopyIn("stoptimes","trip_id", "stop_id", "arrival", "departure", "stop_sequence"))
	if err != nil {
		return
	}

	for _, st := range stopTimes {	
		
		if _, err = stmt.Exec(st.TripID, st.StopID, st.Arrival, st.Departure, st.StopSequence); err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	return tx.Commit()
}


