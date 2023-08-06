package api

import (
	"encoding/json"
	"net/http"
	"henrikkorsgaard.dk/gtfs-service/domain"
	"henrikkorsgaard.dk/gtfs-service/repository"
)

type RestReply struct {
	Stops []domain.Stop `json:"stops"`
}

// API test examples for rest
// Looks like CHI routing framework is better than the fiber framework, which is larger and come with a lot of stuff. 

// Path: "/api/rest/stops"
func StopHandler(w http.ResponseWriter, r *http.Request) {
	repo, err := repository.NewRepository()
	if err != nil {
		panic(err) // we need to handle this later
	}

	stops, err := repo.FetchStops()
	if err != nil {
		panic(err) // we need to handle this later
	}

	reply := RestReply{Stops:stops}

	jsonBytes, err := json.Marshal(reply)
	if err != nil {
		panic(err) // we need to handle this later
	}

	w.Write(jsonBytes)
}
