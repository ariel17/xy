package config

import "flag"

var (
	// DbHost TODO
	DbHost string
	// DbPort TODO
	DbPort int
	// DbName TODO
	DbName string
)

func init() {
	DbHost = "mongo"
	DbPort = 27017
	DbName = "db"
}

// IsTest indicates if the current environment is testing or normal run
func IsTest() bool {
	v := flag.Lookup("test.v")
	return v != nil && v.Value.String() == "true"
}
