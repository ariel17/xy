package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
	"github.com/julienschmidt/httprouter"
)

// GetUsers fetchs an user by indicated ID, if exists.
func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var (
		status int
		result domain.APIResponse
	)

	if id := ps.ByName("id"); id == "" {
		message := "empty id"
		log.Print(message)
		status = http.StatusBadRequest
		result.Message = message

	} else if user, err := dao.Client.GetUser(id); err != nil {
		log.Printf("failed to get user: %v", err)
		status = http.StatusInternalServerError
		result.Message = err.Error()

	} else if user == nil {
		message := fmt.Sprintf("user %s not found", id)
		log.Print(message)
		status = http.StatusNotFound
		result.Message = message

	} else {
		status = http.StatusOK
		result.Success = true
		result.Message = "user found"
		result.Data = user
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

// PostUsers creates a new user with given data.
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
		status = http.StatusCreated
		result.Success = true
		result.Message = "user created"
		result.Data = u
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

// DeleteUsers removes an existing user from storage, if exists.
func DeleteUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var (
		status int
		result domain.APIResponse
	)

	if id := ps.ByName("id"); id == "" {
		message := "empty id"
		log.Print(message)
		status = http.StatusBadRequest
		result.Message = message

	} else if user, err := dao.Client.GetUser(id); err != nil {
		log.Printf("failed to get user: %v", err)
		status = http.StatusInternalServerError
		result.Message = err.Error()

	} else if user == nil {
		message := fmt.Sprintf("user %s not found", id)
		log.Print(message)
		status = http.StatusNotFound
		result.Message = message

	} else if err := dao.Client.DeleteUser(id); err != nil {
		log.Printf("error deleting user %v: %v", user, err)
		status = http.StatusInternalServerError
		result.Message = err.Error()

	} else {
		status = http.StatusOK
		result.Success = true
		result.Message = fmt.Sprintf("user %s deleted", id)
		result.Data = user
		log.Print(user)
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

// PostRegister registers a device into a user's fleet.
func PostRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var (
		status int
		result domain.APIResponse
	)

	// TODO create new pin
	// TODO save new pin
	// TODO include pin value in response

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}
