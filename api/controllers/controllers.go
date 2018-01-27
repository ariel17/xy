package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
)

// Users TODO
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := domain.APIResponse{}
	var status int

	switch r.Method {
	case "GET":

	case "POST":
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var u domain.User
		if err := decoder.Decode(&u); err != nil {
			log.Printf("failed to parse user data: %v", err)
			status = http.StatusForbidden
			result.Message = err.Error()

		} else if err := dao.InsertUser(&u); err != nil {
			log.Printf("failed to save new user: %v", err)
			status = http.StatusInternalServerError
			result.Message = err.Error()

		} else {
			m := "successfully created user"
			log.Printf("%s %v", m, u)
			status = http.StatusCreated
			result.Success = true
			result.Message = m
		}

	default:
		status = http.StatusForbidden
		result.Message = "Not allowed"
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
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
