package storage

import (
	"fmt"

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
	auth := ""
	if s.Auth.User != "" && s.Auth.Password != "" {
		auth = fmt.Sprintf("%s:%s@", s.Auth.User, s.Auth.Password)
	}

	return fmt.Sprintf("%s%s:%d", auth, s.Auth.Host, s.Auth.Port)
}

// Connect TODO
func (s *Storage) Connect() error {

	session, err := mgo.Dial(s.createURL())
	if err != nil {
		return err
	}

	s.session = session
	return nil
}

// Close TODO
func (s *Storage) Close() {
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
