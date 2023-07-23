package main 

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5430
	user     = "henrik"
	password = "henrik"
	dbname   = "gtfs"
	// How do we pull in credentials from somewhere else? Docker file?
  )

func main() {
	fmt.Println("wy")
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println(err)
  		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
  		panic(err)
	}

	// should work now -- so we just need to develop whatever we need!

}