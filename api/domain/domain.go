package domain

import (
	"time"

	"github.com/google/uuid"
)

// User TODO
type User struct {
	ID   int64  `json:"id"`
	Nick string `json:"nick"`
}

// Observable represents a person or object to be tracked.
type Observable struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

// Position represents a point on the Earth surface.
type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ObservablePosition represents a position in Earth for an observable, on a
// given time.
type ObservablePosition struct {
	Observable Observable `json:"observable"`
	Position   Position   `json:"position"`
	CreatedAt  time.Time  `json:"created_at"`
}

// Pin TODO
type Pin struct {
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

// APIResponse TODO
type APIResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	ID      uuid.UUID `json:"id"`
}
