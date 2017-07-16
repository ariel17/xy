package config

import (
	"log"
	"os"
	"strconv"
)

const (
	// EnvDbHost TODO
	EnvDbHost string = "DB_HOST"

	// EnvDbPort TODO
	EnvDbPort string = "DB_PORT"

	// EnvDbDatabase TODO
	EnvDbDatabase string = "DB_DATABASE"

	// EnvDbUser TODO
	EnvDbUser string = "DB_USER"

	// EnvDbPassword TODO
	EnvDbPassword string = "DB_PASSWORD"

	// ErrorInvalidParameters TODO
	ErrorInvalidParameters int = 1

	// ErrorInvalidConfiguration TODO
	ErrorInvalidConfiguration int = 2
)

// Auth TODO
type Auth struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

var (
	// DatabaseAuth TODO
	DatabaseAuth Auth
)

func init() {
	host := os.Getenv(EnvDbHost)

	port, err := strconv.Atoi(os.Getenv(EnvDbPort))
	if err != nil {
		log.Fatalf("Invalid port value for MongoDB: %v", os.Getenv(EnvDbPort))
		os.Exit(ErrorInvalidParameters)
	}

	database := os.Getenv(EnvDbDatabase)
	user := os.Getenv(EnvDbUser)
	password := os.Getenv(EnvDbPassword)

	DatabaseAuth = Auth{
		Host:     host,
		Port:     port,
		Database: database,
		User:     user,
		Password: password,
	}
}
