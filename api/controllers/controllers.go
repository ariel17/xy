package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// GetUsers fetchs an user by indicated ID, if exists.
func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var (
		result domain.APIResponse
	)

	id := ps.ByName("id")
	if id == "" {
		result.Message = "empty id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	hID := bson.ObjectIdHex(id)
	user, err := dao.Client.GetUser(hID)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		result.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
		return
	}

	if user == nil {
		result.Message = fmt.Sprintf("user %s not found", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Success = true
	result.Message = "user found"
	result.Data = user
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// PostUsers creates a new user with given data.
func PostUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	result := domain.APIResponse{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var u domain.User
	if err := decoder.Decode(&u); err != nil {
		result.Message = err.Error()
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(result)
		return
	}

	if err := dao.Client.InsertUser(&u); err != nil {
		result.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Success = true
	result.Message = "user created"
	result.Data = u
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// DeleteUsers removes an existing user from storage, if exists.
func DeleteUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var (
		result domain.APIResponse
	)

	id := ps.ByName("id")
	if id == "" {
		result.Message = "empty id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	hID := bson.ObjectIdHex(id)
	user, err := dao.Client.GetUser(hID)
	if err != nil {
		result.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
		return
	}

	if user == nil {
		result.Message = fmt.Sprintf("user %s not found", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(result)
		return
	}

	if err := dao.Client.DeleteUser(hID); err != nil {
		result.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Success = true
	result.Message = fmt.Sprintf("user %s deleted", id)
	result.Data = user
	w.WriteHeader(http.StatusOK)
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
