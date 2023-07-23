package gtfs

import (
	"context"
	"fmt"
	"os"
	"sync"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

var (
	repoInstance *repository
	repoOnce 	sync.Once
)

// Rewrite to singleton https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
type repository struct {
	pool *pgxpool.Pool
}

func NewRepository() (repo *repository, err error) {

	repoOnce.Do(func(){
		user := os.Getenv("DATABASE_USER")
		pass := os.Getenv("DATABASE_PASS")
		name := os.Getenv("DATABASE_NAME")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")

		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user,pass,host,port,name)
		pool, err := pgxpool.New(context.Background(), connectionString)
		if err != nil {
			// we want to panic here because there is zero chance of recovering from a faulty db config/setup
			panic(err)

		}

		repoInstance = &repository{pool}
	})
	
	return repoInstance, nil
	
}

func (repo *repository) Close(){
	repo.pool.Close()
}

func (repo *repository) IngestStops(stops []Stop) (err error){

	query := `INSERT INTO stops (id, name, description, point) VALUES (@id, @name, @description, @point) ON CONFLICT (id) DO NOTHING`

	batch := &pgx.Batch{}

	for _, s := range stops {
		p := fmt.Sprintf("POINT(%f %f)",s.Lon, s.Lat)

		args := pgx.NamedArgs{
			"id": s.ID,
			"name": s.Name,
			"description":s.Description,
			"point": p,
		}

		batch.Queue(query, args)
	}

	results := repo.pool.SendBatch(context.Background(), batch)

	defer results.Close()
	
	// this is a bit weird, but the exec fetches the result for each query in the queue. We want to do this len(stops) times to make sure no errors happen. 
	for _,_ = range stops {
		_, err = results.Exec()
		if err != nil {
			break 
		}
	}

	if err != nil {
		return
	}

	return results.Close()
	
}





// Not sure I need context
// Read: https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
// Read: https://threedots.tech/post/repository-pattern-in-go/
// Read: https://pkg.go.dev/database/sql#DB.QueryRowContext