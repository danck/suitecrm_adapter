package dto

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	ID        string `json: "id"`
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
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
