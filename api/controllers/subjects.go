package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ariel17/xy/api/domain"
)

// Subjects TODO
func Subjects(w http.ResponseWriter, r *http.Request) {
	subjects := make([]domain.Subject, 0)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(subjects)
	case "POST":
		s := domain.Subject{}
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
