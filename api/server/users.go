package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariel17/xy/api/models"
)

const (
	// ParamNick TODO
	ParamNick string = "nick"
)

// Users TODO
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		r.ParseForm()
		nick := r.Form.Get(ParamNick)
		log.Println("POST fields:", nick)

		u := models.CreateUser(nick)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	default:
		w.WriteHeader(http.StatusForbidden)
	}
}
