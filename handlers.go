package main

import (
	"fmt"
	"net/http"
)

// CustomerHandler
// POST:
//	in:		a CustomerDTO (w/o ID)
//	out:	new customer and response with CustomerDTO+ID
func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//		return CreateCustomer(w, r)
		return
	}

	fmt.Fprint(w, http.StatusNotFound)
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
		//
	}
	if r.Method == "PUT" {
		// TODO(danck) business logic
	}

	fmt.Fprint(w, http.StatusNotFound)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusNotFound)
}
