package dto

import (
	"encoding/json"
	"fmt"
)

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

func (c Customer) JsonString() string {
	cStr, _ := json.Marshal(c)
	return string(cStr[:])
}
