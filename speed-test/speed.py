import csv
from datetime import datetime, timezone, timedelta

start  = datetime.now()

with open("../input/GTFS/stop_times.txt", "r", encoding="utf8") as stfile:
    reader = csv.reader(stfile)
    lines = []
    for row in reader:
        lines.append(row[0])

    print(len(lines))

end = datetime.now()

print((end-start) // timedelta(milliseconds=1))

start  = datetime.now()

# Faster
with open("../input/GTFS/stop_times.txt", "r", encoding="utf8") as stfile:
    lines = stfile.readlines()
    print(len(lines))

end = datetime.now()
print((end-start) // timedelta(milliseconds=1))