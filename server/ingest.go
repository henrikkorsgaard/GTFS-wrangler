package main

import (
	"fmt"
	"henrikkorsgaard.dk/GTFS-wrangler/gtfs"
)

func main(){
	fmt.Println("fmt parking")

	gtfs.DownloadGTFS("https://www.rejseplanen.info/labs/GTFS.zip")
	// we need to parse the GTFS file 
	// then we need to ingest everything. Yeez, I can kill all the process things as well
	// need to convert coordinates to points
	// routes to lines
	

	//need to put on a docker postgis thing here.
}