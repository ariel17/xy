package dao

import (
	"fmt"
	"log"

	"github.com/ariel17/xy/api/config"
	"github.com/ariel17/xy/api/domain"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func init() {
	url := fmt.Sprintf("%s:%s@%s:%d/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	var err error
	if session, err = mgo.Dial(url); err != nil {
		log.Fatal(err)
	}
}

// InsertUser TODO
func InsertUser(u *domain.User) error {
	return getCollection("users").Insert(u)
}

// DeleteUser TODO
func DeleteUser(u *domain.User) error {
	return getCollection("users").Remove(u)
}

func getCollection(collection string) *mgo.Collection {
	return session.DB(config.DbName).C(collection)
}
