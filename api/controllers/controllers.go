package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
	"github.com/julienschmidt/httprouter"
)

// GetUsers TODO
func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	if user, err := dao.Client.GetUser(ps.ByName("id")); err != nil {
		log.Printf("failed to get user: %v", err)
		w.WriteHeader(http.StatusNotFound)
		result := domain.APIResponse{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(result)

	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// PostUsers TODO
func PostUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	result := domain.APIResponse{}
	var status int

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var u domain.User
	if err := decoder.Decode(&u); err != nil {
		log.Printf("failed to parse user data: %v", err)
		status = http.StatusForbidden
		result.Message = err.Error()

	} else if err := dao.Client.InsertUser(&u); err != nil {
		log.Printf("failed to save new user: %v", err)
		status = http.StatusInternalServerError
		result.Message = err.Error()

	} else {
		m := "successfully created user"
		status = http.StatusCreated
		result.Success = true
		result.Message = m
		result.Data = u
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
