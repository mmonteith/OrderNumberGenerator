package clients

import (
	"github.com/stretchr/testify/assert"
	"github.com/globalsign/mgo"

	"testing"
	"errors"
)

func TestCreateMongoClient_OK(t *testing.T) {
	NewMongoSessionProc = func(dInfo *mgo.DialInfo) (*mgo.Session, error) {
		return &mgo.Session{}, nil
	}

	session := CreateMongoSession("localhost:27017")
	if session.Session == nil {
		t.Errorf("Error: Panic occured when trying to create mongo client.")
	}
}

func TestCreateMongoClient_Panic(t *testing.T) {
	NewMongoSessionProc = func(dInfo *mgo.DialInfo) (*mgo.Session, error) {
		return &mgo.Session{}, errors.New("could not connect to server")
	}
	assert.Panics(t, func() { CreateMongoSession("") },
		"Error: Expected panic to occur while creating mongo client but did not.")
}
