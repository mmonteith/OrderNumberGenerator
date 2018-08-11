package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/urbn/ordernumbergenerator/app"
	"github.com/urbn/ordernumbergenerator/app/fixtures"
	"github.com/urbn/ordernumbergenerator/app/mocks"
)

var (
	url          = "/v0/fp-us/sterling-order-number"
	relativePath = "/v0/:siteId/sterling-order-number"

	mockDao = mocks.MockOrdersDao{
		"00000000",
		nil,
	}

	nh = OrderNumberHandler{
		app.PennsylvaniaDC,
		mockDao,
	}
)

func Test_GetOrderNumber_Success(t *testing.T) {
	testRouter := gin.Default()
	testRouter.POST(relativePath, nh.GetOrderNumber)

	response := fixtures.PerformRequest(testRouter, "POST", url)

	if code := response.Code; code != 200 {
		t.Errorf("Handler returned wrong status code: received %v expected %v",
			code, 200)
	}
}

func Test_GetOrderNumber_DataCenterValidation_Error(t *testing.T) {
	handler := nh
	handler.DataCenterId = "US-US"

	testRouter := gin.Default()
	testRouter.POST(relativePath, handler.GetOrderNumber)

	response := fixtures.PerformRequest(testRouter, "POST", url)

	if response.Code != 400 {
		t.Errorf("Handler returned wrong status code: received %v expected %v",
			response.Code, 400)
	}
}

func Test_GetOrderNumber_BrandIdValidation_Error(t *testing.T) {
	invalidPath := "/v0/no-brand/sterling-order-number"

	testRouter := gin.Default()
	testRouter.POST(relativePath, nh.GetOrderNumber)

	response := fixtures.PerformRequest(testRouter, "POST", invalidPath)

	if response.Code != 400 {
		t.Errorf("Handler returned wrong status code: received %v expected %v",
			response.Code, 400)
	}
}

func Test_GetOrderNumber_OrderNumDao_Error(t *testing.T) {
	handler := nh
	handler.Dao = mocks.MockOrdersDao{
		OrderNum: "0",
		Error:    &fixtures.MockOrderDaoError,
	}

	testRouter := gin.Default()
	testRouter.POST(relativePath, handler.GetOrderNumber)

	response := fixtures.PerformRequest(testRouter, "POST", url)

	if response.Code != 400 {
		t.Errorf("Handler returned wrong status code: received %v expected %v",
			response.Code, 400)
	}
}

func Test_GetOrderNumber_OrderNumValidation_Error(t *testing.T) {
	handler := nh
	handler.Dao = mocks.MockOrdersDao{
		OrderNum: "0",
		Error:    nil,
	}

	testRouter := gin.Default()
	testRouter.POST(relativePath, handler.GetOrderNumber)

	response := fixtures.PerformRequest(testRouter, "POST", url)

	if response.Code != 500 {
		t.Errorf("Handler returned wrong status code: received %v expected %v",
			response.Code, 500)
	}
}

func Test_buildOrderNumber(t *testing.T) {
	orderNumber := buildOrderNumber("US-PA", "uo", "00000001")
	if orderNumber != "TP00000001" {
		t.Errorf("Exepected orderNumber TP00000001 but got %s", orderNumber)
	}
}

func Test_validateBrandAndDataCenter_Success(t *testing.T) {
	result := validateBrandAndDataCenter("uo-us", app.PennsylvaniaDC)
	if result != nil {
		t.Errorf("Expected an error status of 200 but got %d", result.Code)
	}
}

func Test_validateBrandAndDataCenter_Invalid_BrandID_Error(t *testing.T) {
	result := validateBrandAndDataCenter("foo", app.PennsylvaniaDC)
	if result.Code != 400 {
		t.Errorf("Expected an error status of 400 but got %d", result.Code)
	}
}

func Test_validateBrandAndDataCenter_Invalid_DataCenterID_Error(t *testing.T) {
	result := validateBrandAndDataCenter("uo-us", "foo")
	if result.Code != 400 {
		t.Errorf("Expected an error status of 400 but got %d", result.Code)
	}
}

func Test_validateOrderNumber_Success(t *testing.T) {
	result := validateOrderNumber("00000001")
	if result != nil {
		t.Errorf("Expected an error status of 200 but got %d", result.Code)
	}
}

func Test_validateOrderNumber_Invalid_OrderNumber_Len_Error(t *testing.T) {
	result := validateOrderNumber("000")
	if result.Code != 500 {
		t.Errorf("Expected an error status of 500 but got %d", result.Code)
	}
}

func Test_validateOrderNumber_Invalid_OrderNumber_Error(t *testing.T) {
	result := validateOrderNumber("!!!")
	if result.Code != 500 {
		t.Errorf("Expected an error status of 500 but got %d", result.Code)
	}
}
