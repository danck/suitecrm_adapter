package main

import (
	"fmt"
	"github.com/danck/hawai-suitecrm/customer"
	"github.com/danck/hawai-suitecrm/order"
	"net/http"
)

// CustomersHandler
// POST:
//  Creates a customer in SuiteCRM. The ID is set by the CRM.
//	in:		a CustomerDTO (w/o ID)
//	out:	new customer and response with CustomerDTO+ID
func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		customer, err := customer.New("dummy name", "dummy address")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(customer.JsonString())
		fmt.Fprintf(w, customer.JsonString())
		return
	}

	http.Error(w,
		http.StatusText(http.StatusNotImplemented),
		http.StatusNotImplemented)
}

// OrdersHandler
// POST		Creates an order for a given customer
//	in:		OrderDTO, CustomerID
//	out:	OrderDTO
// PUT		Signals that an order has been paid
//	in:		OrderID, AdressDTO
//	out:
func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		order, err := order.New("dummy customer id for order")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(order.JsonString())
		return
	}
	if r.Method == "PUT" {
		// TODO(danck) business logic
	}

	http.Error(w,
		http.StatusText(http.StatusNotImplemented),
		http.StatusNotImplemented)
}

// DefaultHandler
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w,
		http.StatusText(http.StatusNotImplemented),
		http.StatusNotImplemented)
}
