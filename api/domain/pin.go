package domain

import (
	"time"

	"github.com/ariel17/xy/api/config"
)

// Pin TODO
type Pin struct {
	Value     string
	CreatedAt time.Time
}

// IsValid TODO
func (p Pin) IsValid() bool {
	return time.Since(p.CreatedAt) < config.PINDuration
}

// NewPin TODO
func NewPin() *Pin {
	return &Pin{
		Value:     "abc123",
		CreatedAt: time.Now(),
	}
}
