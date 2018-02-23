package domain

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User TODO
type User struct {
	ID   bson.ObjectId `json:"_id" bson:"_id"`
	Nick string        `json:"nick" bson:"name"`
}

// Observable represents a thing (person or object) to be tracked.
type Observable struct {
	ID     bson.ObjectId `json:"_id" bson:"_id"`
	UserID int64         `json:"user_id" bson:"user_id"`
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
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	UserID    string        `json:"user_id"`
	Value     string        `json:"value"`
	CreatedAt time.Time     `json:"created_at"`
}

// APIResponse TODO
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
