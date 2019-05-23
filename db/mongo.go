package db

import (
	"github.com/globalsign/mgo"
)

const (
	dbURL = "mongodb://localhost:27017/?readPreference=primary"
)

// Dial ...
func Dial() *mgo.Session {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session
}

// Collection ...
func Collection(collection string) *mgo.Collection {
	session := Dial()
	c := session.DB("bus").C(collection)
	return c
}
