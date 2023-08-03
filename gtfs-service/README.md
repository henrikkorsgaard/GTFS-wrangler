# GTFS service

This is a small micro service that will do the following:

- Ingest GTFS files from provided URL when the source has changed (HTTP statuscode 304)
- Persist GTFS data into a PostgreSQL database with PostGIS extension
- Coming: Provide the following API interfaces: RESTful, GraphQL, SQL (odata?) for consumption.

I'm developing this because I need a GTFS microservice for Danish GTFS data for specific projects. With that goal in mind, I prioritize features that are required by the other projects. 

## A note on the tests and test data
I try to add tests for each part of the codebase as I implement the features. The tests are not designed to or have perfect coverage (e.g. database fetch does not check values, just that the number of rows match). I use them to guide my implementation and check that new stuff does not break old stuff. 

I use a small subset of the Danish Rejseplan GTFS dataset as test data. Currently, the test data does not include all GTFS files or fields. The following fields are missing in the test data and therefor not implemented:

- frequencies.txt: Is empty in the Danish GTFS data. 
- fare_attributes.txt: Is not part of the data set.
- fare_rules.txt: Is not part of the data set.
- pathways.txt: Is not part of the data set.
- levels.txt: Is not part of the data set.
- feed_info.txt: Is not part of the data set.
- translations.txt: Is not part of the data set.

The plan is to create test data that will provide full test coverage of the GTFS specification down the road. 

## Supported GTFS fields

See: https://developers.google.com/transit/gtfs/reference#field_definitions

Currently the following fields are supported:

- Agency: https://developers.google.com/transit/gtfs/reference#agencytxt
- Stops: https://developers.google.com/transit/gtfs/reference#stopstxt
- Routes: https://developers.google.com/transit/gtfs/reference#routestxt
- Trips: https://developers.google.com/transit/gtfs/reference#tripstxt
- Stop times: https://developers.google.com/transit/gtfs/reference#stop_timestxt
- Calendar: https://developers.google.com/transit/gtfs/reference#calendartxt
- Calendar dates: https://developers.google.com/transit/gtfs/reference#calendar_datestxt
- Shapes: https://developers.google.com/transit/gtfs/reference#shapestxt


## Documentation
TODO

## Other libraries
There are a couple of useful GTFS libraries and implementations out there. I have use a couple of them for reference and learned a lot from their implementation and coverage. 

#### [GO GTFS](https://github.com/artonge/go-gtfs):
This Golang library loads GTFS files from a directory into one GFTS struct. It is currently in maintenance mode. I have modified parts of the [mapToDestination](https://github.com/artonge/go-csv-tag/blob/4b40f225e91a009021bac2ae6fd04a3d90c58b12/load.go#L142) function and use the [storeValue](https://github.com/artonge/go-csv-tag/blob/4b40f225e91a009021bac2ae6fd04a3d90c58b12/load.go#L194) function in my code. 

This project is licensed under GPL-3.0 because I have taken inspiration from and use code from this library.

#### [GTFS via postgres](https://github.com/public-transport/gtfs-via-postgres)
This NodeJS cli tool loads GTFS files from a directory into a PostgreSQL database. You can query the data using SQL and/or (with the right config) use it in connection with [PostGraphile](https://www.graphile.org/postgraphile/) to get a GraphQL interface (neat!).

#### [GTFS to Geojson](https://github.com/BlinkTagInc/gtfs-to-geojson)
This NodeJS cli tool/library can convert GTFS data into geojson. 
