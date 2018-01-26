package config

import (
	"time"
)

const (
	// ErrorInvalidParameters TODO
	ErrorInvalidParameters = iota

	// ErrorInvalidConfiguration TODO
	ErrorInvalidConfiguration

	// PINDuration TODO
	PINDuration = 3 * time.Hour

	// PINMaxAmount TODO
	PINMaxAmount = 10
)

var (
	// DbHost TODO
	DbHost string
	// DbPort TODO
	DbPort int
	// DbName TODO
	DbName string
	// DbUser TODO
	DbUser string
	// DbPassword TODO
	DbPassword string
)

func init() {
	DbHost = "localhost"
	DbPort = 9090
	DbName = "db"
	DbUser = "username"
	DbPassword = "password"
}
