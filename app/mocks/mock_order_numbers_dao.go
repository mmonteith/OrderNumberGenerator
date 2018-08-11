package mocks

import "github.com/urbn/ordernumbergenerator/app"

type MockOrdersDao struct {
	OrderNum string
	Error    *app.Error
}

func (od MockOrdersDao) GetOrderNumberByBrandAndDataCenter(brand, dataCenter string) (string, *app.Error) {
	return od.OrderNum, od.Error
}
