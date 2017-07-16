package server

import (
	"encoding/json"
	"net/http"

	"github.com/ariel17/xy/api/models"
)

// Users TODO
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		u := models.CreateUser()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	default:
		w.WriteHeader(http.StatusForbidden)
	}
}
