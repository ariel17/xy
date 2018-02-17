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
	DeleteUser(u *domain.User) error
	GetUser(id string) (*domain.User, error)
}

var db DB

func init() {
	if config.IsTest() {
		db = &MockDB{}
	} else {
		db = &MongoDB{}
	}

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
}
