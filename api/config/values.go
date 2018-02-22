package config

import "os"

const (
	// ProductionEnv TODO
	ProductionEnv = "production"
)

var (
	// DbHost TODO
	DbHost string
	// DbPort TODO
	DbPort int
	// DbName TODO
	DbName string
	// Environment TODO
	Environment string
)

func init() {
	DbHost = "mongo"
	DbPort = 27017
	DbName = "db"
	Environment = os.Getenv("ENVIRONMENT")
}
