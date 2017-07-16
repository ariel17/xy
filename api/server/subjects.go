package server

import (
	"encoding/json"
	"net/http"
	"time"
)

// Position Earth coordinates given by GPS device, asociated to a point i
// time.
type Position struct {
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
}

// Subject Represents a person or object to be tracked.
type Subject struct {
	Current Position
	History []Position
}

// Subjects TODO
func Subjects(w http.ResponseWriter, r *http.Request) {
	subjects := make([]Subject, 0)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(subjects)
	case "POST":
		s := Subject{}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)
	case "PUT":
		// TODO Update an existing record.
	case "DELETE":
		// TODO Remove the record.
	default:
		// TODO Give an error message.
	}
}
