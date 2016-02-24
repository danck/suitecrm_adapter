package dto

import (
	"strings"
	"testing"
)

func testCustomerJsonString(t *testing.T, title string) {
	expectedString := `{"Id":"123","FirstName":"testFirstName","LastName":"testLastName"}`
	customer := Customer{
		ID:        "123",
		FirstName: "testFirstName",
		LastName:  "testLastName",
		Email:     "test@mail.com",
		Password:  "testPassword",
		//Address
		//ShipmentAddress
		//InvoiceAddress
	}
	c := customer.JsonString()
	if strings.Compare(c, expectedString) != 0 {
		t.Errorf("Expected: %s\nGot: %s", expectedString, c)
	}
}
