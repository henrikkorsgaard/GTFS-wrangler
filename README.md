## thank you svelte for overriding my Readme!


## TODO:
### Server:
- Implement all the GTFS file loaders
-- Routes
-- Then stop_times for progres benchmark
-- TODO: handle types in unmarshal
- Implement a chan solutionf for giving progress feedback to front end
- Implement a gtfs specific error handling for better errors?
## Notes on ingesting GTFS csv to structs
After a couple of experiments, I have decided not to use the [existing package](https://github.com/artonge/go-gtfs/tree/master) that can marshal into GTFS. First, it uses an underlying [csv](https://github.com/artonge/go-csv-tag/tree/master) that is significantly slower (factor 2 - 3) than writing the csv-to-gtfs conversion by hand. Second, I need some additional control over the tags (e.g. implementing require, optional and conditional optional checks). Third, I need to inject the 


## Application ideas

### Widgets
Tailored feed that just tells when the next bus departs from specific stop
"Get me out out here" Autocomplete + stop selection = show a way out from X + destination
"Take me to X" Autocomplet + stop selection = show how to get to X soonest / fastest

#### Some benchmarks
When I parse the stop_times.txt file with approx 4.085.000 rows, then I get the following times
- GO-CSV-TAG: 8.8s
- Handrolled: 4.5s

I have not implemented additional lookup in the GO-CSV-TAG method, but would likely add 2s+. When I do pprof I can see that reflect structTag lookups take approx 2s for one structTag.

I have not implemented "required" fields in any of the above. I expect it to have a marginal impact on the handrolled method versus the csv-reflect tag lookup.