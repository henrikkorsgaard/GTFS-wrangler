drop table if exists stops;
drop table if exists trips;
drop table if exists routes;
drop table if exists shapes;
drop table if exists stoptimes;

CREATE TABLE IF NOT EXISTS stops (
  id varchar primary key,
  name varchar not null,
  description varchar,
  geo_point geography(point, 4326) not null,
  parent_station varchar not null
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

