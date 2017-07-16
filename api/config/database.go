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

	// EnvDbName TODO
	EnvDbName string = "DB_NAME"

	// EnvDbUser TODO
	EnvDbUser string = "DB_USER"

	// EnvDbPassword TODO
	EnvDbPassword string = "DB_PASSWORD"
)

// Auth TODO
type Auth struct {
	Host     string
	Port     int
	Name     string
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

	name := os.Getenv(EnvDbName)
	user := os.Getenv(EnvDbUser)
	password := os.Getenv(EnvDbPassword)

	DatabaseAuth = Auth{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     user,
		Password: password,
	}
}
