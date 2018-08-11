package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/urbn/ordernumbergenerator/app"
)

type OrderNumberHandler struct {
	DataCenterId string
	Dao          app.OrderNumberDao
}

func (nh OrderNumberHandler) GetOrderNumber(c *gin.Context) {
	siteId := c.Param("siteId")

	error := validateBrandAndDataCenter(siteId, nh.DataCenterId)
	if error != nil {
		fmt.Printf("Order number handler called with error: %s", error.Message)
		c.JSON(error.Code, gin.H{"status": error.Status, "message": error.Message})
		return
	}

	number, error := nh.Dao.GetOrderNumberByBrandAndDataCenter(app.BrandId[siteId], nh.DataCenterId)
	if error != nil {
		c.JSON(error.Code, gin.H{"status": error.Status, "message": error.Message})
		return
	}

	error = validateOrderNumber(number)
	if error != nil {
		fmt.Printf("Sterling Order Number created was invalid: %s", number)
		c.JSON(error.Code, gin.H{"status": error.Status, "message": error.Message})
		return
	}

	order := app.SterlingOrderNumberResponse{
		Brand:               app.BrandId[siteId],
		DataCenterId:        nh.DataCenterId,
		SterlingOrderNumber: buildOrderNumber(nh.DataCenterId, app.BrandId[siteId], number),
	}
	fmt.Printf("Order number handler called and responded with order number %s using brand %s and data center %s", order.SterlingOrderNumber, order.Brand, order.DataCenterId)
	c.JSON(200, order)
}

func buildOrderNumber(dataCenter string, brandId string, orderNum string) string {
	sterlingOrderNumber := fmt.Sprintf("%v%v%v", app.BrandPrefix[brandId], app.DataCenterPrefix[dataCenter],
		orderNum)
	return sterlingOrderNumber
}

func validateBrandAndDataCenter(siteId, dataCenter string) *app.Error {
	if _, ok := app.BrandId[siteId]; ok {
	} else {
		return &app.Error{
			Code:    400,
			Status:  http.StatusBadRequest,
			Message: "Invalid site ID."}
	}

	if _, ok := app.DataCenterId[dataCenter]; ok {
	} else {
		return &app.Error{
			Code:    400,
			Status:  http.StatusBadRequest,
			Message: "Invalid data center ID."}
	}
	return nil
}

//TODO @(aholsinger@urbn.com) A15-14308 Rewrite to validate whole order number and not just DAO's returned integer
func validateOrderNumber(orderNumber string) *app.Error {
	if _, err := strconv.Atoi(orderNumber); err != nil {
		return &app.Error{
			Code:    500,
			Status:  http.StatusInternalServerError,
			Message: "Invalid Sterling Order Number response"}
	}

	if len(orderNumber) != 8 {
		return &app.Error{
			Code:    500,
			Status:  http.StatusInternalServerError,
			Message: "Invalid Sterling Order Number response"}
	}
	return nil
}
