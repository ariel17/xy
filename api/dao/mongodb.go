package dao

import (
	"fmt"

	"github.com/ariel17/xy/api/config"
	"github.com/ariel17/xy/api/domain"
	"gopkg.in/mgo.v2"
)

// MongoDB TODO
type MongoDB struct {
	session *mgo.Session
}

func (m *MongoDB) createURL() string {
	auth := fmt.Sprintf("%s:%s@", config.DbUser, config.DbPassword)
	database := fmt.Sprintf("/%s", config.DbName)
	return fmt.Sprintf("%s%s:%d%s", auth, config.DbHost, config.DbPort, database)
}

// Connect TODO
func (m *MongoDB) Connect() error {
	session, err := mgo.Dial(m.createURL())
	if err != nil {
		return err
	}
	m.session = session
	return err
}

// Close TODO
func (m *MongoDB) Close() error {
	m.session.Close()
	return nil
}

// InsertUser TODO
func (m *MongoDB) InsertUser(u *domain.User) error {
	return m.getCollection("users").Insert(m)
}

// DeleteUser TODO
func (m *MongoDB) DeleteUser(u *domain.User) error {
	return m.getCollection("users").Remove(m)
}

func (m *MongoDB) getCollection(collection string) *mgo.Collection {
	return m.session.DB(config.DbName).C(collection)
}
