drop table if exists stops;
drop table if exists trips;
drop table if exists routes;
drop table if exists shapes;

CREATE TABLE IF NOT EXISTS stops (
  id varchar primary key,
  name varchar not null,
  description varchar,
  point geography(point, 4326) not null,
  parentstation varchar not null
);

CREATE TABLE IF NOT EXISTS routes (
  id varchar primary key, 
  agencyid varchar not null,
  name varchar not null,
  longname varchar not null,
  type varchar not null
);

CREATE TABLE IF NOT EXISTS trips (
  id integer primary key, 
  serviceid varchar not null, 
  routeid varchar not null,
  shapeid integer not null,
  tripheadsign varchar
);

CREATE TABLE IF NOT EXISTS shapes (
  id integer primary key,
  line geography(linestring,4326) not null
);

