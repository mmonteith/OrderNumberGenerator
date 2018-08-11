package daos

import (
	"errors"
	"testing"

	"github.com/urbn/ordernumbergenerator/app/mocks"
)

func TestCreateOrderNumberDao(t *testing.T) {
	dao := CreateOrderNumberDao(mocks.MockSession{}, mocks.MockCollection{})

	if dao.MongoSession == nil {
		t.Error("Error: nil Session field in OrderNumberDao")
	}
	if dao.MongoCollection == nil {
		t.Error("Error: nil Collection field in OrderNumberDao")
	}
}

func TestOrderNumberDao_GetOrderNumberByBrandAndDataCenter_Success(t *testing.T) {
	session := mocks.NewMockSession()
	collection := session.DB("OrderNumberGen").C("ordernums")

	dao := CreateOrderNumberDao(session, collection)

	orderNum, err := dao.GetOrderNumberByBrandAndDataCenter("an", "US-NV")

	if orderNum != "00000001" {
		t.Errorf("Error: Incorrect ordernumber returned: expected \"00000001\" but got \"%s\"", orderNum)
	}
	if err != nil {
		t.Errorf("Error: Unexpected error occured: %s", err.Message)
	}
}

func TestOrderNumberDao_GetOrderNumberByBrandAndDataCenter_Error(t *testing.T) {
	session := mocks.NewMockSession()
	collection := session.DB("OrderNumberGen").C("ordernums")

	mocks.Error = errors.New("apply error")
	dao := CreateOrderNumberDao(session, collection)

	_, err := dao.GetOrderNumberByBrandAndDataCenter("bad brand", "US-NV")

	if err != nil && err.Code != 500 {
		t.Errorf("Expected an error code of 400 but got %d", err.Code)
	}
}
