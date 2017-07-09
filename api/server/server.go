package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	SUBJECTS_PATH string = "/subjects"
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

// Configure Maps URL paths into handlers.
func Configure() {
	http.HandleFunc(SUBJECTS_PATH, subjects)
}

// Start Servers API endpoint in indicated ip and port.
func Start(ip string, port string) error {
	address := fmt.Sprintf("%s:%s", ip, port)
	log.Println("Starting server in", address)
	return http.ListenAndServe(address, nil)
}

func subjects(w http.ResponseWriter, r *http.Request) {
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
