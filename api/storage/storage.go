package storage

import (
	"fmt"

	"log"

	"github.com/ariel17/xy/api/config"
	"gopkg.in/mgo.v2"
)

var (
	instance *Storage
)

// Storage TODO
type Storage struct {
	Auth    config.Auth
	session *mgo.Session
}

func (s Storage) createURL() string {
	auth, database := "", ""
	if s.Auth.User != "" && s.Auth.Password != "" {
		auth = fmt.Sprintf("%s:%s@", s.Auth.User, s.Auth.Password)
	}

	if s.Auth.Name != "" {
		database = fmt.Sprintf("/%s", s.Auth.Name)
	}

	return fmt.Sprintf("%s%s:%d%s", auth, s.Auth.Host, s.Auth.Port, database)
}

// Connect TODO
func (s *Storage) Connect() error {

	session, err := mgo.Dial(s.createURL())
	if err != nil {
		return err
	}

	s.session = session
	log.Println("Successfully connected to MongoDB :)")
	return nil
}

// Close TODO
func (s *Storage) Close() {
	log.Println("Closing MongoDB connections.")
	s.session.Close()
}

// CreateStorage TODO
func CreateStorage() (*Storage, error) {
	if instance == nil {
		instance = &Storage{Auth: config.DatabaseAuth}
		return instance, instance.Connect()
	}

	return instance, nil
}
