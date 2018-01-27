package config

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
