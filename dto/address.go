package dto

import (
	"encoding/json"
)

type Address struct {
	Id          string `json:"Id"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Street      string `json:"Street"`
	HouseNumber string `json:"HouseNumber"`
	Zipcode     string `json:"Zipcode"`
	City        string `json:"City"`
	Country     string `json:"Country"`
}

func (a *Address) JsonString() string {
	aStr, _ := json.Marshal(a)
	return string(aStr[:])
}
