drop table if exists agency;
drop table if exists stops;
drop table if exists trips;
drop table if exists routes;
drop table if exists shapes;
drop table if exists stoptimes;

CREATE TABLE IF NOT EXISTS agency (
  agency_id TEXT PRIMARY KEY,
  agency_name TEXT NOT NULL,
  agency_url TEXT NOT NULL, 
  agency_timezone TEXT NOT NUll,
  /* https://github.com/public-transport/gtfs-via-postgres/blob/main/lib/agency.js 
    Suggest there might be an issue with timezones that needs to be accounted for, e.g. when the timezone of the data does not match the timezone of the database, then there will be a conflict.
  */
  agency_lang TEXT,
  agency_phone TEXT,
  agency_fare_url TEXT,
  agency_email TEXT
);

CREATE TABLE IF NOT EXISTS stops (
  id TEXT PRIMARY KEY,
  stop_code TEXT,
  stop_name TEXT NOT NULL,
  stop_desc TEXT,
  stop_loc geography(point, 4326) NOT NULL,
  zone_id TEXT,
  stop_url TEXT,
  location_type TEXT,
  /* we can use the same enum technique from https://github.com/public-transport/gtfs-via-postgres/blob/main/lib/stops.js */
  parent_station TEXT NOT NULL,
  stop_timezone TEXT, 
  wheelchair_boarding TEXT,
  level_id TEXT,
  platform_code TEXT
);

CREATE TABLE IF NOT EXISTS routes (
  id varchar primary key, 
  agency_id varchar not null,
  short_name varchar not null,
  long_name varchar not null,
  type varchar not null
);

CREATE TABLE IF NOT EXISTS trips (
  id integer primary key, 
  service_id varchar not null, 
  route_id varchar not null,
  shape_id integer not null,
  trip_headsign varchar
);

CREATE TABLE IF NOT EXISTS shapes (
  id integer primary key,
  geo_line geography(linestring,4326) not null
);

CREATE TABLE IF NOT EXISTS stoptimes (
  trip_id integer,
  stop_id varchar,
  arrival time not null,
  departure time not null,
  stop_sequence varchar not null,
  PRIMARY KEY (trip_id, stop_id)
);

