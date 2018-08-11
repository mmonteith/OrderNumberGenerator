package mocks

import (
	"github.com/globalsign/mgo"
	"github.com/urbn/ordernumbergenerator/app"
	"github.com/urbn/ordernumbergenerator/app/fixtures"
)

var Error error = nil

type MockSession struct{}

func (ms MockSession) Close() {}

func NewMockSession() MockSession {
	return MockSession{}
}

func (ms MockSession) DB(name string) app.DataLayer {
	mockDatabase := MockDatabase{}
	return mockDatabase
}

type MockCollection struct{}

func (mc MockCollection) Find(query interface{}) app.Query {
	return MockQuery{}
}

type MockDatabase struct{}

func (md MockDatabase) C(name string) app.Collection {
	mockCollection := MockCollection{}
	return mockCollection
}

type MockQuery struct{}

func (mmc MockQuery) Apply(change mgo.Change, result interface{}) (*mgo.ChangeInfo, error) {
	if Error != nil{
		return &mgo.ChangeInfo{}, Error
	}

	*result.(*app.MongoDocument) = fixtures.ApplyResultAN
	return &mgo.ChangeInfo{}, nil
}
