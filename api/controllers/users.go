package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariel17/xy/api/domain"
)

// Users TODO
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		u := domain.NewUser("ariel")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)

	case "POST":
		r.ParseForm()
		nick := r.Form.Get("nick")
		log.Println("POST fields:", nick)

		u := domain.NewUser(nick)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)

	default:
		w.WriteHeader(http.StatusForbidden)
	}
}
