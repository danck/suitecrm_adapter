package hawaicrm

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CustomersHandler
// POST:
//  Creates a customer in SuiteCRM. The ID is set by the CRM.
// 	in:		a CustomerDTO (w/o ID)
//	out:	new customer and response with CustomerDTO+ID
func CustomersHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return createCustomer(w, r)
	}

	return NotImplemented{"Prefix exists, but no matching HTTP method"}
}

func createCustomer(w http.ResponseWriter, r *http.Request) error {
	log.Println("Creating customer")
	var c Customer
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	if err != nil {
		log.Println("Decode failed", err)
		return BadRequest{err.Error()}
	}

	//log.Println("Decoded request body: ", c.String())
	nameValueList := []KeyValuePair{
		KeyValuePair{"name", c.FirstName + " " + c.LastName},
	}

	resp, err := SetEntry("Accounts", nameValueList)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, resp.(string))
	//log.Println("customer:", customer.JsonString())
	//fmt.Fprintf(w, customer.JsonString())
	return nil
}

// OrdersHandler
// POST		Creates an order for a given customer
//	in:		OrderDTO, CustomerID
//	out:	OrderDTO (200 OK)
// PUT		Signals that an order has been paid
//	in:		OrderID, AdressDTO (200 OK)
//	out:	(200 OK)
func OrdersHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return createOrder(w, r)
	}
	if r.Method == "PUT" {
		return markAsPaid(w, r)
	}
	return NotImplemented{"Prefix exists, but no matching HTTP method"}
}

// markAsPaid requires an order/invoice id and sets its status to 'Paid'
func markAsPaid(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// createOrder creates a new order with the given data and associates it with
// an existing customer
func createOrder(w http.ResponseWriter, r *http.Request) error {
	order, err := NewOrder("dummy customer id for order")
	if err != nil {
		return BadRequest{err.Error()}
	}
	log.Println(order.JsonString())
	resp, err := SetEntry("AOS_Invoices", nil)
	if resp == nil {
		return nil
	}
	return nil
}

func defaultHandler(w http.ResponseWriter, r *http.Request) error {
	return NotImplemented{"Prefix doesn't exist"}
}
