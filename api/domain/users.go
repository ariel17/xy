package models

import (
	"time"

	"github.com/ariel17/xy/api/config"
)

const (
	// ModelName TODO
	ModelName string = "users"
)

// PIN TODO
type Pin struct {
	Value     string
	CreatedAt time.Time
}

// IsValid TODO
func (p Pin) IsValid() bool {
	elapsed := time.Now().Sub(p.CreatedAt)
	return elapsed.Hours() <= config.PINDuration
}

// CreatePIN TODO
func CreatePin() *Pin {
	return &Pin{
		Value:     "abc123",
		CreatedAt: time.Now(),
	}
}

// User TODO
type User struct {
	Nick        string `json:"nick"`
	PendingPins []Pin  `json:"pending_pins"`
}

// CreateUser TODO
func CreateUser(nick string) User {
	// TODO save user into storage
	return User{
		Nick:        nick,
		PendingPins: make([]Pin, config.PinMaxAmount),
	}
}

func (u *User) Save() error {
	return nil
}

func (u *User) Delete() error {
	return nil
}

func (u *User) GetModelName() string {
	return ModelName
}
