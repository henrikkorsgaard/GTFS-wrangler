import csv
import json

# I just need one trip -> route + times

route_id = "23982_3"
shape_id = "3142"
trip_id = "111718624"
stops = []
stop_ids = {}

trip_row = {}
route_row = {}
shape_rows = []

with open("input/GTFS/routes.txt", "r", encoding="utf8") as tfile:
    reader = csv.reader(tfile)
    for row in reader:
        if row[0] == route_id:
            route_row = row

with open("input/GTFS/trips.txt", "r", encoding="utf8") as tfile:
    reader = csv.reader(tfile)
    for row in reader:
        if row[2] == trip_id:
            trip_row = row

with open("input/GTFS/shapes.txt", "r", encoding="utf8") as tfile:
    reader = csv.reader(tfile)
    for row in reader:
        if row[0] == shape_id:
            shape_rows.append(row)

with open("input/GTFS/stop_times.txt", "r", encoding="utf8") as stfile:
    reader = csv.reader(stfile)

    for row in reader:
        if row[0] == trip_id:
            """ """
            stop_ids[row[3]] = ""
            stops.append(
                {"stop_id": row[3], "arrival": row[1], "name": "", "lat": "", "lng": ""}
            )


with open("input/GTFS/stops.txt", "r", encoding="utf8") as sfile:
    reader = csv.reader(sfile)

    for row in reader:
        if row[0] in stop_ids:
            for stop in stops:
                if row[0] == stop["stop_id"]:
                    stop["name"] = row[2]
                    stop["lat"] = row[4]
                    stop["lng"] = row[5]




features = []
coordinates = []
latlngs = []

for row in shape_rows: 
    latlngs.append([row[1],row[2]])
    coordinates.append([row[2],row[1]]) # Because GeoJson uses lng lat

properties = {
    "type": "transit_route",
    "agency_name": "Midttrafik",
    "agency_id": "281",
    "route_id": route_id,
    "route_short_name": route_row[2],
    "route_type": route_row[4],
    "shape_id": shape_id,
    "trip_headsign": trip_row[3],
    "direction_id": trip_row[5],
    "latlngs": latlngs,
    "stops": stops,
}

geometry = {
        "type": "LineString",
        "coordinates": coordinates
    }

route_feature = {"type": "Feature", "properties": properties, "geometry": geometry}

features.append(route_feature)

for s in stops:
    sprops = {
        "type": "transit_stop",
        "name": s["name"]
    }

    sgeo = {
        "type": "point",
        "coordinates": [s["lng"], s["lat"]]
    }

    sfeat = {"type": "Feature", "properties": sprops, "geometry": sgeo}
    features.append(sfeat)

jsdata = {
    "type": "FeatureCollection",
    "features": features,
}

js = json.dumps(jsdata, indent=4, ensure_ascii=False)
name = "281_" + route_id + "_" + route_row[2] + "_" + trip_id + ".geojson"

with open("trip-output/"+name, "w") as outfile:
    outfile.write(js)
    