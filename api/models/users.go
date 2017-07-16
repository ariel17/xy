package models

import (
	"time"

	"github.com/ariel17/xy/api/config"
)

// PIN TODO
type PIN struct {
	Value     string
	CreatedAt time.Time
}

// IsValid TODO
func (p PIN) IsValid() bool {
	elapsed := time.Now().Sub(p.CreatedAt)
	return elapsed.Hours() <= float64(config.PINValidHours)
}

// CreatePIN TODO
func CreatePIN() PIN {
	return PIN{
		Value:     createPINValue(),
		CreatedAt: time.Now(),
	}
}

func createPINValue() string {
	// TODO
	return "abc123"
}

// User TODO
type User struct {
	PendingPIN []PIN
}

// CreateUser TODO
func CreateUser() User {
	// TODO save user into storage
	return User{
		PendingPIN: make([]PIN, config.PINMaxAmount),
	}
}
