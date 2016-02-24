package dto

import (
	"strings"
	"testing"
)

func testAddressJsonString(t *testing.T, title string) {
	expectedString := `{"Id":"123","FirstName":"testFirstName","LastName":"testLastName","Street":"Test Street","HouseNumber":"13a"}`
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
	a := address.JsonString()
	if strings.Compare(a, expectedString) != 0 {
		t.Errorf("Expected: %s\nGot: %s", expectedString, a)
	}
}
