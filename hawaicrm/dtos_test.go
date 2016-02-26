package hawaicrm

import (
	"strings"
	"testing"
)

func TestAddressJsonString(t *testing.T) {
	address := Address{
		Id:          "123",
		FirstName:   "testFirstName",
		LastName:    "testLastName",
		Street:      "Test Street",
		HouseNumber: "13a",
		//Zipcode
		//City
		//Country
	}
	expectedString := `{"Id":"123","FirstName":"testFirstName","LastName":"testLastName","Street":"Test Street","HouseNumber":"13a","Zipcode":"","City":"","Country":""}`
	a := address.JsonString()
	if strings.Compare(a, expectedString) != 0 {
		t.Errorf("Expected: %s\nGot: %s", expectedString, a)
	}
}

func TestCustomerJsonString(t *testing.T) {
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
	expectedString := `{"Id":"123","FirstName":"testFirstName","LastName":"testLastName","Email":"test@mail.com","Password":"testPassword","Address":{"Id":"","FirstName":"","LastName":"","Street":"","HouseNumber":"","Zipcode":"","City":"","Country":""},"ShipmentAddress":{"Id":"","FirstName":"","LastName":"","Street":"","HouseNumber":"","Zipcode":"","City":"","Country":""},"InvoiceAddress":{"Id":"","FirstName":"","LastName":"","Street":"","HouseNumber":"","Zipcode":"","City":"","Country":""}}`
	c := customer.JsonString()
	if strings.Compare(c, expectedString) != 0 {
		t.Errorf("Expected: %s\nGot: %s", expectedString, c)
	}
}
