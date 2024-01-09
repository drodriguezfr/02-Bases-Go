package repository

import (
	"app/internal"
	"errors"
	"fmt"
)

var (
	ErrorInvalidCustomer = errors.New("invalid field value")
)

func NewCustomersMap() *CustomersMap {
	return &CustomersMap{customers: make(map[int]internal.Customer)}
}

type CustomersMap struct {
	customers map[int]internal.Customer
}

func (c *CustomersMap) Save(customer *internal.Customer) (err error) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error: client already exists")
		}
	}()
	_, id := c.customers[customer.Id]
	if id {
		panic("User already exist")
	}

	// auto-increment id
	//(*c).lastId++
	// set id
	//customer.Id = (*c).lastId

	// save
	c.customers[customer.Id] = *customer
	return

}

func (c *CustomersMap) ValidateZero(customer *internal.Customer) (string, err error) {
	print(customer.File)
	if customer.File == "0" || customer.Home == "0" || customer.Name == "0" || customer.Phone == "0" || customer.Id == 0 {
		err = ErrorInvalidCustomer
		return
	}
	return
}
