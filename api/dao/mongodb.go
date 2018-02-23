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
func (m *MongoDB) DeleteUser(id string) error {
	return m.getCollection("users").RemoveId(id)
}

// GetUser TODO
func (m *MongoDB) GetUser(id string) (*domain.User, error) {
	var u domain.User
	err := m.getCollection("users").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
	return &u, err
}

// GetPendingPin returns a registered pin for indicated ID, or error.
func (m *MongoDB) GetPendingPin(id string) (*domain.Pin, error) {
	var pin domain.Pin
	err := m.getCollection("pending_pins").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&pin)
	return &pin, err
}

// InsertPendingPin saves a new pin for an indicated user.
func (m *MongoDB) InsertPendingPin(p *domain.Pin) error {
	p.ID = bson.NewObjectId()
	return m.getCollection("pending_pins").Insert(p)
}

// DeletePendingPin removes a pending pin by its ID.
func (m *MongoDB) DeletePendingPin(id string) error {
	return m.getCollection("pending_pins").RemoveId(id)
}

func (m *MongoDB) getCollection(collection string) *mgo.Collection {
	return m.session.DB(config.DbName).C(collection)
}
