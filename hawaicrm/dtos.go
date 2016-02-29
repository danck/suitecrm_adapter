package hawaicrm

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	ID          string `json:"Id"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Street      string `json:"Street"`
	HouseNumber string `json:"HouseNumber"`
	Zipcode     string `json:"Zipcode"`
	City        string `json:"City"`
	Country     string `json:"Country"`
}

func (a *Address) JSONString() string {
	aStr, _ := json.Marshal(a)
	return string(aStr[:])
}

type Customer struct {
	ID              string  `json:"Id"`
	FirstName       string  `json:"FirstName"`
	LastName        string  `json:"LastName"`
	Email           string  `json:"Email"`
	Password        string  `json:"Password"`
	Address         Address `json:"Address"`
	ShipmentAddress Address `json:"ShipmentAddress"`
	InvoiceAddress  Address `json:"InvoiceAddress"`
}

func (c Customer) String() string {
	return fmt.Sprintf("Customer: (ID: %s, Name: %s, Address: %s)", c.ID, c.FirstName, c.LastName)
}

func (c Customer) JSONString() string {
	cStr, _ := json.Marshal(c)
	return string(cStr[:])
}

type Order struct {
	OrderID    string
	CustomerID string
}

func (o Order) JSONString() string {
	oStr, _ := json.Marshal(o)
	return string(oStr[:])
}

func NewOrder(customerID string) (*Order, error) {
	if customerID == "" {
		return nil, fmt.Errorf("empty name")
	}
	return &Order{"", customerID}, nil
}
