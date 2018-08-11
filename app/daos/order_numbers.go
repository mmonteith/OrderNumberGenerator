package daos

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"

	"github.com/urbn/ordernumbergenerator/app"
)

type OrderNumberDao struct {
	MongoSession    app.Session
	MongoCollection app.Collection
}

func CreateOrderNumberDao(mongoSession app.Session, mongoCollection app.Collection) OrderNumberDao {
	ond := OrderNumberDao{
		MongoSession:    mongoSession,
		MongoCollection: mongoCollection,
	}
	return ond
}

func (ond OrderNumberDao) GetOrderNumberByBrandAndDataCenter(brand, dataCenter string) (string, *app.Error) {
	document := app.MongoDocument{}
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"orderNumber": 1}},
		ReturnNew: false,
	}

	//TODO @(aholsinger@urbn.com) A15-14308 Split this up to separately check for connection errors and improper number format errors
	_, err := ond.MongoCollection.Find(bson.M{"brandId": brand, "dataCenterId": dataCenter}).Apply(change, &document)
	if err != nil{
		fmt.Printf("Unable to retrieve number from Mongo Collection: %s", ond.MongoCollection)
		return "", &app.Error{
			Code: 500,
			Status: http.StatusInternalServerError,
			Message: "Unable to retrieve number from MongoDB"}
	}

	return fmt.Sprintf("%08d", document.OrderNumber), nil
}
