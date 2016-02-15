package main

import (
	"encoding/json"
	"fmt"
	"github.com/danck/hawai-suitecrm/dto"
	"github.com/danck/hawai-suitecrm/errtypes"
	"net/http"
)

// CustomersHandler
// POST:
//  Creates a customer in SuiteCRM. The ID is set by the CRM.
// 	in:		a CustomerDTO (w/o ID)
//	out:	new customer and response with CustomerDTO+ID
func CustomersHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		var c dto.Customer
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&c)
		if err != nil {
			return errtypes.BadRequest{err.Error()}
		}

		customer, err := dto.NewCustomer("dummy name", "dummy address")
		if err != nil {
			return err
		}
		fmt.Println(customer.JsonString())
		fmt.Fprintf(w, customer.JsonString())
		return nil
	}

	return errtypes.NotImplemented{"Prefix exists, but no matching HTTP method"}
}

// OrdersHandler
// POST		Creates an order for a given customer
//	in:		OrderDTO, CustomerID
//	out:	OrderDTO
// PUT		Signals that an order has been paid
//	in:		OrderID, AdressDTO
//	out:
func OrdersHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		order, err := dto.NewOrder("dummy customer id for order")
		if err != nil {
			return errtypes.BadRequest{err.Error()}
		}
		fmt.Println(order.JsonString())
		return nil
	}
	if r.Method == "PUT" {
		// TODO(danck) business logic
	}

	return errtypes.NotImplemented{"Prefix exists, but no matching HTTP method"}
}

// DefaultHandler
func DefaultHandler(w http.ResponseWriter, r *http.Request) error {
	return errtypes.NotImplemented{"No matching prefix"}
}
