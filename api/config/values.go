package config

import (
	"log"
	"os"
	"strconv"
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
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
)

func init() {
	DbHost := os.Getenv("DB_HOST")

	DbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("invalid port value for MongoDB", err)
		os.Exit(ErrorInvalidParameters)
	}

	DbName := os.Getenv("DB_NAME")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
}
