package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ariel17/xy/api/domain"
)

// Users TODO
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// TODO get saved user

	case "POST":
		// TODO create user and save

	default:
		w.WriteHeader(http.StatusForbidden)
		result := domain.APIResponse{
			Message: "Not allowed.",
		}
		json.NewEncoder(w).Encode(result)
	}

}

// Register TODO
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case "POST":
		// TODO

	case "OPTIONS":
		// TODO

	default:
		w.WriteHeader(http.StatusBadRequest)
		result := domain.APIResponse{
			Message: "Not allowed.",
		}
		json.NewEncoder(w).Encode(result)
	}
}
