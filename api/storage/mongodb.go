package storage

import (
	"fmt"

	"log"

	"github.com/ariel17/xy/api/config"
	"github.com/ariel17/xy/api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB TODO
type MongoDB struct {
	Auth    config.Auth
	session *mgo.Session
}

func (m *MongoDB) createURL() string {
	auth, database := "", ""
	if m.Auth.User != "" && m.Auth.Password != "" {
		auth = fmt.Sprintf("%s:%s@", m.Auth.User, m.Auth.Password)
	}

	if m.Auth.Name != "" {
		database = fmt.Sprintf("/%s", m.Auth.Name)
	}

	return fmt.Sprintf("%s%s:%d%s", auth, m.Auth.Host, m.Auth.Port, database)
}

// Connect TODO
func (m *MongoDB) Connect() error {

	session, err := mgo.Dial(m.createURL())
	if err != nil {
		return err
	}

	m.session = session
	log.Println("Successfully connected to MongoDB :)")
	return nil
}

// Close TODO
func (m *MongoDB) Close() error {
	log.Println("Closing MongoDB connections.")
	m.session.Close()
	return nil
}

// Insert TODO
func (mdb *MongoDB) Insert(m *models.Model) error {
	c := mdb.getCollection(m)
	return c.Insert(m)
}

// Delete TODO
func (mdb *MongoDB) Delete(m *models.Model) error {
	c := mdb.getCollection(m)
	return c.Remove(m)
}

// Get TODO
func (mdb *MongoDB) Get() *models.User {
	u := models.User{}
	c := mdb.session.DB(mdb.Auth.Name).C("user")
	c.Find(bson.M{"nick": "ariel17"}).One(&u)
	return &u
}

func (mdb *MongoDB) getCollection(m *models.Model) *mgo.Collection {
	return mdb.session.DB(mdb.Auth.Name).C((*m).GetModelName())
}
