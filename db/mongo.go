package db

import (
	"time"

	"github.com/globalsign/mgo"
)

const (
	dbURL = "mongodb://localhost:27017/bus"
)

// GlobalMgoSession ...
var GlobalMgoSession *mgo.Session

func init() {
	globalMgoSession, err := mgo.DialWithTimeout(dbURL, 10*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
}

// CloneSession ...
func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}
