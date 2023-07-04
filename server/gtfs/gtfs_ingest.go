package gtfs

import (
	"time"
	"net/http"
	"io"
)

/**
* I want this to assume that it downloads a zip file as a byte stream
* then it threats is as an unzip file
* then we ingest the various structs in a handheld manner
* -> we get rows and decide what to do with them
* -> parse accordingly, e.g. shapes need to be handled differently
* -> bulk inserted into the db
* TODO: handle http statuscodes for logging...
* 
*/

// how do we maintain date -- we ask on the day before!

func DownloadGTFS(url string) (bytes []byte, err error){
	
	n := time.Now()
	// need to be 24 hours when live
	day, err := time.ParseDuration("24h")
	if err != nil {
		return
	}
	ts := n.Add(-day).Format(time.RFC1123)

	client := &http.Client{
		Timeout:10*time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Add("If-Modified-Since", ts)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// we kinda want to handle a few cases here, but so far we hope nobody moves the files... 
	if resp.StatusCode != http.StatusOK {
		return
	}
	bytes, err = io.ReadAll(resp.Body)
	
	return 
}