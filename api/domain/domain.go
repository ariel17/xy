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

// NewUser creates a new user with setted ID.
func NewUser(nick string) *User {
	return &User{
		ID:   bson.NewObjectId(),
		Nick: nick,
	}
}

// Device represents a thing being tracked.
type Device struct {
	ID     bson.ObjectId `json:"_id" bson:"_id"`
	UserID bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Model  string        `json:"model" bson:"model"`
}

// NewDevice creates a new device for given user ID with setted ID, empty model.
func NewDevice(userID bson.ObjectId) *Device {
	return &Device{
		ID:     bson.NewObjectId(),
		UserID: userID,
	}
}

// Position represents a point on the Earth surface.
type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// DeviceRegistration is the data required to associate a device with an user.
type DeviceRegistration struct {
	UserID bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Pin    string        `json:"pin"`
	Device Device        `json:"device"`
}

// DevicePosition represents a position in Earth for a device, on a given time.
type DevicePosition struct {
	Device    Device    `json:"device"`
	Position  Position  `json:"position"`
	CreatedAt time.Time `json:"created_at"`
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
