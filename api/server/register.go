package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const (
	PIN_NAME string = "pin"
)

// Result The operation result information and flags.
type Result struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	ID      *uuid.UUID `json:"id,omitempty"`
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case "POST":
		r.ParseForm()
		pin := r.Form.Get(PIN_NAME)
		// TODO save pin to storage
		// TODO associate pin to some user
		log.Println("Gained PIN code:", pin)
		w.WriteHeader(http.StatusCreated)
		id := uuid.New()
		result := Result{true, "Successfully created ID.", &id}
		log.Println(fmt.Sprintf("Result: %v", result))
		json.NewEncoder(w).Encode(result)

	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		result := Result{true, "Allowed methods: POST", nil}
		json.NewEncoder(w).Encode(result)

	default:
		w.WriteHeader(http.StatusBadRequest)
		result := Result{false, "Not allowed.", nil}
		json.NewEncoder(w).Encode(result)
	}
}
