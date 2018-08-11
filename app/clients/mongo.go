package clients

import (
	"time"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/urbn/ordernumbergenerator/app"
	"fmt"
)

var (
	NewMongoSessionProc = NewSession
)

func NewSession(dialInfo *mgo.DialInfo) (*mgo.Session, error) { return mgo.DialWithInfo(dialInfo) }

func CreateMongoSession(mongoServers string) MongoSession {
	mhosts := strings.Split(mongoServers, ",")
	dialInfo := &mgo.DialInfo{
		Addrs:   mhosts,
		Timeout: 10 * time.Second,
	}
	mongoSession, err := NewMongoSessionProc(dialInfo)
	if err != nil {
		fmt.Printf("Unable to connect to Mongo session at address: %s", dialInfo.Addrs)
		panic(err)
	}
	mongoSession.SetMode(mgo.Monotonic, true)
	return MongoSession{mongoSession}
}

type MongoSession struct {
	*mgo.Session
}

func (ms MongoSession) DB(name string) app.DataLayer {
	return &MongoDatabase{Database: ms.Session.DB(name)}
}

type MongoDatabase struct {
	*mgo.Database
}

func (md MongoDatabase) C(name string) app.Collection {
	return MongoCollection{Collection: md.Database.C(name)}
}

type MongoQuery struct {
	*mgo.Query
}

func (mq MongoQuery) Apply(change mgo.Change, result interface{}) (*mgo.ChangeInfo, error) {
	return mq.Query.Apply(change, result)
}

type MongoCollection struct {
	*mgo.Collection
}

func (mc MongoCollection) Find(query interface{}) app.Query {
	return MongoQuery{
		Query: mc.Collection.Find(query),
	}
}
