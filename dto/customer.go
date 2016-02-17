package dto

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	ID        string
	FirstName string
	//LastName       string
	//AccountName    string
	//EmailAddress   string
	PrimaryAddress string
}

func (c Customer) String() string {
	return fmt.Sprintf("Customer: (ID: %s, Name: %s, Address: %s)", c.ID, c.FirstName, c.PrimaryAddress)
}

func (c Customer) JsonString() string {
	cStr, _ := json.Marshal(c)
	return string(cStr[:])
}

func NewCustomer(name string, address string) (*Customer, error) {
	if name == "" {
		return nil, fmt.Errorf("empty name")
	}
	return &Customer{"", name, address}, nil
}
