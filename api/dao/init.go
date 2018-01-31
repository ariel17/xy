package dao

import (
	"github.com/ariel17/xy/api/config"
	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

func init() {
	if config.IsTest() {
	} else {
		newRealMongoDB()
	}
}
