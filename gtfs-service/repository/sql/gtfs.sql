drop table if exists agency;
drop table if exists stops;
drop table if exists trips;
drop table if exists routes;
drop table if exists shapes;
drop table if exists stoptimes;

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

