package dao

import (
	"fmt"

	"github.com/ariel17/xy/api/config"
	"github.com/ariel17/xy/api/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB is the implementation for the DB abstraction.
type MongoDB struct {
	session *mgo.Session
}

// Connect opens the connection to database.
func (m *MongoDB) Connect() error {
	url := fmt.Sprintf("%s:%d", config.DbHost, config.DbPort)
	var err error
	m.session, err = mgo.Dial(url)
	return err
}

// InsertUser TODO
func (m *MongoDB) InsertUser(u *domain.User) error {
	u.ID = bson.NewObjectId()
	return m.getCollection("users").Insert(u)
}

// DeleteUser TODO
func (m *MongoDB) DeleteUser(u *domain.User) error {
	return m.getCollection("users").Remove(u)
}

// GetUser TODO
func (m *MongoDB) GetUser(id string) (*domain.User, error) {
	var u domain.User
	err := m.getCollection("users").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
	return &u, err
}

func (m *MongoDB) getCollection(collection string) *mgo.Collection {
	return m.session.DB(config.DbName).C(collection)
}
