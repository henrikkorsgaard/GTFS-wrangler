drop table if exists stops;

CREATE TABLE IF NOT EXISTS stops (
  id serial primary key,
  name varchar not null,
  description varchar,
  point geography(point, 4326) not null
);
