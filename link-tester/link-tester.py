# https://www.rejseplanen.info/labs/GTFS.zip

import requests
from datetime import timedelta, datetime

then = datetime.now() - timedelta(seconds=100) # We need to store the last download date.
# Run as a script that gulps and insert into the database.
# Then we have a database with whatever we need.
# And we can construct the stuff we need from the data.
ts = then.strftime('%a, %d %b %Y %H:%M:%S GMT')
print(ts)
headers = {"If-Modified-Since": ts}
x = requests.get('https://api.statbank.dk/v1/subjects', headers=headers)
print(x.status_code)

# What are the scenarios
## Updated -> gulp 
## Not updated -> do nothing
## Any other update than 200 / 304 means I need to do some work. 
## Single point of failure
## Who implements this
#### Rejseplanen yes
#### DST no
