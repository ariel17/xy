package dao

import (
	"log"

	"github.com/ariel17/xy/api/config"
	"github.com/ariel17/xy/api/domain"
)

// DB is the abstraction for the underlying storage implementation.
type DB interface {
	Connect() error
	InsertUser(u *domain.User) error
	DeleteUser(id string) error
	GetUser(id string) (*domain.User, error)
}

// Client holds the implementation instance, based on the environment.
var Client DB

func init() {
	if config.Environment == config.ProductionEnv {
		Client = &MongoDB{}
	} else {
		Client = &MockDB{}
	}

	if err := Client.Connect(); err != nil {
		log.Fatal(err)
	}
}
