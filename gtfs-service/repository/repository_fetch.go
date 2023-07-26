package repository

import (
	"henrikkorsgaard.dk/gtfs-service/domain"
	"github.com/twpayne/go-geom/encoding/ewkb"
)

// TODO: https://pkg.go.dev/github.com/twpayne/go-geom
// https://github.com/twpayne/go-geom/blob/master/examples/postgis/main.go
func (repo *repository) FetchStops() (stops []domain.Stop, err error){
	rows, err := repo.db.Query("SELECT id, name, description, parent_station, ST_AsBinary(geo_point) FROM stops;")
	defer rows.Close()
	
	for rows.Next() {
		s := domain.Stop{}
		var p ewkb.Point
		err = rows.Scan(&s.ID, &s.Name, &s.Description,&s.ParentStation, &p)
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