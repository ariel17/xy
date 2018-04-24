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

// Users -----------------------------------------------------------------------

// InsertUser TODO
func (m *MongoDB) InsertUser(u *domain.User) error {
	u.ID = bson.NewObjectId()
	return m.getCollection("users").Insert(u)
}

// DeleteUser TODO
func (m *MongoDB) DeleteUser(id bson.ObjectId) error {
	return m.getCollection("users").RemoveId(id)
}

// GetUser TODO
func (m *MongoDB) GetUser(id bson.ObjectId) (*domain.User, error) {
	var u domain.User
	query := bson.M{"_id": id}
	err := m.getCollection("users").Find(query).One(&u)
	return &u, err
}

// GetPendingPin returns a registered pin for indicated ID, or error.
func (m *MongoDB) GetPendingPin(id bson.ObjectId) (*domain.Pin, error) {
	var pin domain.Pin
	query := bson.M{"_id": id}
	err := m.getCollection("pending_pins").Find(query).One(&pin)
	return &pin, err
}

// InsertPendingPin saves a new pin for an indicated user.
func (m *MongoDB) InsertPendingPin(p *domain.Pin) error {
	p.ID = bson.NewObjectId()
	return m.getCollection("pending_pins").Insert(p)
}

// DeletePendingPin removes a pending pin by its ID.
func (m *MongoDB) DeletePendingPin(id bson.ObjectId) error {
	return m.getCollection("pending_pins").RemoveId(id)
}

// Users + Devices -------------------------------------------------------------

// GetUserDevices returns all devices registered for indicated user and/or
// errors.
func (m *MongoDB) GetUserDevices(id bson.ObjectId) ([]domain.Device, error) {
	devices := []domain.Device{}
	query := bson.M{"user_id": id}
	err := m.getCollection("devices").Find(query).All(devices)
	return devices, err
}

// Devices ---------------------------------------------------------------------

// GetDevice returns a device matching the indicated ID or a possible error.
func (m *MongoDB) GetDevice(id bson.ObjectId) (*domain.Device, error) {
	var d domain.Device
	query := bson.M{"_id": id}
	err := m.getCollection("devices").Find(query).One(&d)
	return &d, err
}

// privates --------------------------------------------------------------------

func (m *MongoDB) getCollection(collection string) *mgo.Collection {
	return m.session.DB(config.DbName).C(collection)
}
