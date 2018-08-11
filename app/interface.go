package app

import (
	"github.com/globalsign/mgo"
)

// Encoder - Interface to encoding/json for unit test reasons
type Encoder interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

// Session - Interface for mongo Session
type Session interface {
	DB(string) DataLayer
}

// DataLayer - Interface for mongo Database
type DataLayer interface {
	C(string) Collection
}

// Query - Interface for mongo Query
type Query interface {
	Apply(mgo.Change, interface{}) (*mgo.ChangeInfo, error)
}

// Collection - Interface for mongo Collection
type Collection interface {
	Find(interface{}) Query
}

// OrderNumberDao - Interface for accessing mongo
type OrderNumberDao interface {
	GetOrderNumberByBrandAndDataCenter(string, string) (string, *Error)
}
