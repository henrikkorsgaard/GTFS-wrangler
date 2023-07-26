package repository

import (
	
	"fmt"
	"os"
	"sync"

	"database/sql"

	_ "github.com/lib/pq"

)

var (
	repoInstance *repository
	repoOnce 	sync.Once
)

// Rewrite to singleton https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
type repository struct {
	db *sql.DB
}

func NewRepository() (repo *repository, err error) {

	repoOnce.Do(func(){
		user := os.Getenv("DATABASE_USER")
		pass := os.Getenv("DATABASE_PASS")
		name := os.Getenv("DATABASE_NAME")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")

		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user,pass,host,port,name)
		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			// we want to panic here because there is zero chance of recovering from a faulty db config/setup
			panic(err)

		}

		repoInstance = &repository{db}
	})
	
	return repoInstance, nil
	
}

func (repo *repository) Close(){
	repo.db.Close()
}

