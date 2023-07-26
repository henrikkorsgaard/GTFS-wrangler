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
	/*
	query := `INSERT INTO stops (id, name, description, geo_point, parent_station) VALUES (@id, @name, @description, @point, @parentstation) ON CONFLICT (id) DO NOTHING`

	batch := &pgx.Batch{}

	for _, s := range stops {
		p := wkb.Point{Point:&s.GeoPoint}
		fmt.Printf("%#v\n", p)
		args := pgx.NamedArgs{
			"id": s.ID,
			"name": s.Name,
			"description":s.Description,
			"point": p,
			"parentstation": s.ParentStation,
		}

		batch.Queue(query, args)
	}

	results := repo.pool.SendBatch(context.Background(), batch)

	defer results.Close()
	
	// The batch processing in PGX is a bit weird in terms of design. The exec fetches the result for each query in the queue for each item in the queue. The easiest way to fetch all is to use the length of the queue.

	for i, n := 0, batch.Len(); i < n ; i++ {
		_, err = results.Exec()
		if err != nil {
			break 
		}
	} 


	if err != nil {
		return
	}

	return results.Close()
	*/
	return
}

func (repo *repository) IngestRoutes(routes []domain.Route) (err error){
	/*
	query := `INSERT INTO routes (id, agency_id, short_name, long_name, type) VALUES (@id, @agencyid, @name, @longname, @type) ON CONFLICT (id) DO NOTHING`

	batch := &pgx.Batch{}

	for _, r := range routes {
		
		args := pgx.NamedArgs{
			"id": r.ID,
			"agencyid": r.AgencyID,
			"name":r.Name,
			"longname": r.LongName,
			"type": r.Type,
		}

		batch.Queue(query, args)
	}

	results := repo.pool.SendBatch(context.Background(), batch)

	defer results.Close()
	 
	for i, n := 0, batch.Len(); i < n ; i++ {
		_, err = results.Exec()
		if err != nil {
			break 
		}
	}

	if err != nil {
		return
	}

	return results.Close()
	*/
	return
}

func (repo *repository) IngestTrips(trips []domain.Trip) (err error){
	/*
	query := `INSERT INTO trips (id, service_id, route_id,shape_id, trip_headsign) VALUES (@id, @serviceid, @routeid,@shapeid,@tripheadsign) ON CONFLICT (id) DO NOTHING`

	batch := &pgx.Batch{}

	for _, t := range trips {
		args := pgx.NamedArgs{
			"id": t.ID,
			"serviceid": t.ServiceID,
			"routeid":t.RouteID,
			"shapeid": t.ShapeID,
			"tripheadsign": t.TripHeadsign,
		}

		batch.Queue(query, args)
	}

	results := repo.pool.SendBatch(context.Background(), batch)

	defer results.Close()
	
	for i, n := 0, batch.Len(); i < n ; i++ {
		_, err = results.Exec()
		if err != nil {
			break 
		}
	} 

	if err != nil {
		return
	}

	return results.Close()
	*/
	return 
}

func (repo *repository) IngestShapes(shapes []domain.Shape) (err error){
	/*

	query := `INSERT INTO shapes (id, geo_line) VALUES (@id, @line) ON CONFLICT (id) DO NOTHING`
	
	batch := &pgx.Batch{}

	for _, s := range shapes {

		coordStrings := make([]string, 0)
		for _, ll := range s.Coordinates {
			llstr := fmt.Sprintf("%f %f", ll.Longitude, ll.Latitude)
			coordStrings = append(coordStrings, llstr)
		}
	
		line := fmt.Sprintf("LINESTRING(%s)", strings.Join(coordStrings[0:], ","))

		args := pgx.NamedArgs{
			"id": s.ID,
			"line": line,
		}

		batch.Queue(query, args)
	}

	results := repo.pool.SendBatch(context.Background(), batch)

	defer results.Close()
	
	for i, n := 0, batch.Len(); i < n ; i++ {
		_, err = results.Exec()
		if err != nil {
			break 
		}
	} 

	if err != nil {
		return
	}

	return results.Close()
	*/ 
	return
}

func (repo *repository) IngestStopTimes(stopTimes []domain.StopTime) (err error){
	/*
	query := `INSERT INTO stoptimes (trip_id, stop_id, arrival, departure, stop_sequence) VALUES (@tripid, @stopid, @arrival, @departure, @sequence) ON CONFLICT (trip_id, stop_id) DO NOTHING`

	batch := &pgx.Batch{}

	for _, s := range stopTimes {

		args := pgx.NamedArgs{
			"tripid": s.TripID,
			"stopid": s.StopID,
			"arrival":s.Arrival,
			"departure": s.Departure,
			"sequence": s.StopSequence,
		}

		batch.Queue(query, args)
		
	}

	results := repo.pool.SendBatch(context.Background(), batch)

	defer results.Close()
	
	for i, n := 0, batch.Len(); i < n ; i++ {
		_, err = results.Exec()
		if err != nil {
			break 
		}
	} 

	if err != nil {
		return
	}

	return results.Close()
	*/
	return
}
