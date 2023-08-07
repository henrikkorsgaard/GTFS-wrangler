package api


import (
	"fmt"
	"testing"
	"net/http"
	"io"
	"net/http/httptest"
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/twpayne/go-geom"
	"github.com/joho/godotenv"

	"henrikkorsgaard.dk/gtfs-service/domain"
	"henrikkorsgaard.dk/gtfs-service/testutils"
	"henrikkorsgaard.dk/gtfs-service/repository"
)

func init(){
	fmt.Println("Running rest tests")
	godotenv.Load("../config_dev.env")
	err := testutils.ResetDatabase("../repository/sql/gtfs.sql")
	if err != nil {
		panic(err)
	}
}

/*
	- I expect it to return an array of json stops
	- Where does the data come from?
	-- I could make a struct with a stop, insert it into the db and the query that -> unmarshal and deep compare?

	-- I could load some test data like I do in the repository test?
	-- Via the real data (some setup)
	-- or via a sql file I query
*/

func TestStopHandler(t *testing.T) {

	// Setting up the test data
	point := geom.NewPoint(geom.XY)
	point.SetSRID(4326)

	point.MustSetCoords(geom.Coord{10.345034934460,55.354917787122})

	gpoint := domain.JSONGeoPoint{*point}

	testStop := domain.Stop{ID:"000461011300", Name:"Dyrupgårds Alle (Odense Kommune)",Lat: 0, Lon: 0, GeoPoint:gpoint}

	stops := []domain.Stop{testStop}
	repo, err := repository.NewRepository()
	err = repo.IngestStops(stops)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	ts := httptest.NewServer(http.HandlerFunc(StopHandler))
	defer ts.Close()

	url := fmt.Sprintf("%s/api/rest/stops", ts.URL)

	res, err := ts.Client().Get(url)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	reply, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	replyStops := RestReply{}

	err = json.Unmarshal(reply, &replyStops)
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	

	assert.Equal(t, testStop, replyStops.Stops[0])
}