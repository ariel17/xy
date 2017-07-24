package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariel17/xy/api/config"
	"github.com/ariel17/xy/api/models"
	"github.com/ariel17/xy/api/storage"
)

const (
	// ParamNick TODO
	ParamNick string = "nick"
)

// Users TODO
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		db := storage.MongoDB{
			Auth: config.DatabaseAuth,
		}
		u := db.Get()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)

	case "POST":
		r.ParseForm()
		nick := r.Form.Get(ParamNick)
		log.Println("POST fields:", nick)

		u := models.CreateUser(nick)
		u.Save()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	default:
		w.WriteHeader(http.StatusForbidden)
	}
}
