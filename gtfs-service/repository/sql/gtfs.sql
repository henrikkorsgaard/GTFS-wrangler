drop table if exists agency;
drop table if exists stops;
drop table if exists trips;
drop table if exists routes;
drop table if exists shapes;
drop table if exists stoptimes;
drop table if exists calendar;

CREATE TABLE IF NOT EXISTS agency (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  url TEXT NOT NULL, 
  timezone TEXT NOT NUll,
  /* https://github.com/public-transport/gtfs-via-postgres/blob/main/lib/agency.js 
    Suggest there might be an issue with timezones that needs to be accounted for, e.g. when the timezone of the data does not match the timezone of the database, then there will be a conflict.
  */
  lang TEXT,
  phone TEXT,
  fare_url TEXT,
  email TEXT
);

CREATE TABLE IF NOT EXISTS stops (
  id TEXT PRIMARY KEY,
  code TEXT,
  name TEXT NOT NULL,
  description TEXT,
  location geography(point, 4326) NOT NULL,
  zone_id TEXT,
  url TEXT,
  location_type TEXT,
  /* we can use the same enum technique from https://github.com/public-transport/gtfs-via-postgres/blob/main/lib/stops.js */
  parent_station TEXT NOT NULL,
  timezone TEXT, 
  wheelchair_boarding TEXT,
  level_id TEXT,
  platform_code TEXT
);

CREATE TABLE IF NOT EXISTS routes (
  id TEXT PRIMARY KEY, 
  agency_id TEXT NOT NULL,
  short_name TEXT NOT NULL,
  long_name TEXT NOT NULL,
  description TEXT,
  type TEXT NOT NULL,
  url TEXT,
  color TEXT,
  text_color TEXT,
  sort_order INTEGER,
  continuous_pickup TEXT,
  continuous_drop_off TEXT
);

CREATE TABLE IF NOT EXISTS trips (
  id TEXT PRIMARY KEY, 
  route_id TEXT NOT NULL,
  service_id TEXT NOT NULL,  
  shape_id TEXT NOT NULL, 
  headsign TEXT,
  name TEXT,
  block_id TEXT,
  wheelchair_accessible TEXT,
  bikes_allowed TEXT
);

CREATE TABLE IF NOT EXISTS stoptimes (
  trip_id TEXT,
  stop_id TEXT,
  arrival TIME NOT NULL,
  departure TIME NOT NULL,
  stop_sequence TEXT NOT NULL,
  stop_headsign TEXT,
  pickup_type TEXT,
  drop_off_type TEXT,
  continuous_pickup TEXT,
  continuous_drop_off TEXT,
  shape_dist_traveled TEXT,
  timepoint TEXT,
  PRIMARY KEY (trip_id, stop_id)
);

CREATE TABLE IF NOT EXISTS calendar (
  service_id TEXT PRIMARY KEY, /*this is the foreign key as well -- check how to do this */
  monday BOOLEAN NOT NULL, /* could be should be enum? Only have 2 states 1: service available for day: 0: service unavailable for day, see https://developers.google.com/transit/gtfs/reference#calendartxt*/
  tuesday BOOLEAN NOT NULL,
  wednesday BOOLEAN NOT NULL,
  thursday BOOLEAN NOT NULL,
  friday BOOLEAN NOT NULL,
  saturday BOOLEAN NOT NULL,
  sunday BOOLEAN NOT NULL,
  start_date DATE NOT NULL,
  end_date DATE NOT NULL
);

/* Hertil */

CREATE TABLE IF NOT EXISTS shapes (
  id integer primary key,
  geo_line geography(linestring,4326) not null
);



