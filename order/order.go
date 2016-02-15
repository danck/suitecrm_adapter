package order

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	OrderID    string
	CustomerID string
}

func (o Order) JsonString() string {
	oStr, _ := json.Marshal(o)
	return string(oStr[:])
}

func New(customerId string) (*Order, error) {
	if customerId == "" {
		return nil, fmt.Errorf("empty name")
	}
	return &Order{"", customerId}, nil
}
