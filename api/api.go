package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Position struct {
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
}

type Subject struct {
	Current Position
	History []Position
}

func subjects(w http.ResponseWriter, r *http.Request) {
	subjects := make([]Subject, 0)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(subjects)
	case "POST":
		s := Subject{}
		json.NewEncoder(w).Encode(s)
	case "PUT":
		// Update an existing record.
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}

func main() {
	http.HandleFunc("/subjects", subjects)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Args[1]), nil))
}
