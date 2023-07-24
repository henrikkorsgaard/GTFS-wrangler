package main

import (
	"os"
	"fmt"
	"henrikkorsgaard.dk/gtfs-service/ingest"
)

func main(){
	fmt.Println("fmt parking")
	url := "https://www.rejseplanen.info/labs/GTFS.zip"

	// This needs to come in from somewhere
	bytes, err := ingest.Download(url)
	if err != nil {
		panic(err)
	}

	// what is a good ingest patter?
	// download, then convert to structs, then insert into the db?
	// or just download, the put into DB?
	
	// ALSO: do we have 


	// if returning zero bytes, we assume the resource has not been modified.
	if len(bytes) == 0 {
		fmt.Printf("Nothing downloaded from '%s'. Resource not modified since last time.\n", url)
		os.Exit(0)
	}

	gtfsFiles, err := ingest.ParseZipIntoFiles(bytes)
	if err != nil {
		panic(err)
	}

	for _,f := range gtfsFiles {
		
		// now we can look at how to ingest this shit.
		// do we marshal into a struct or what?

		// we use the existing unmarshal functions, but handle it a bit more elegantly this time. 
		// some should be marshalled with a specific focus, e.g. shapes
		// and include geocoordinates

		// so I guess its a file per file case of implementation
		
		fmt.Println(f.Name)
		break
	}
	
	// we need to parse the GTFS file 
	// then we need to ingest everything. Yeez, I can kill all the process things as well
	// need to convert coordinates to points
	// routes to lines
	

	//need to put on a docker postgis thing here.
}