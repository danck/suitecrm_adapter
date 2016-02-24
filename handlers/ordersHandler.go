package handlers

import (
	con "gitlab.com/danck/hawai-suitecrm/connector"
	"gitlab.com/danck/hawai-suitecrm/dto"
	"gitlab.com/danck/hawai-suitecrm/errtypes"
	"log"
	"net/http"
)

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
	return errtypes.NotImplemented{"Prefix exists, but no matching HTTP method"}
}

// markAsPaid requires an order/invoice id and sets its status to 'Paid'
func markAsPaid(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// createOrder creates a new order with the given data and associates it with
// an existing customer
func createOrder(w http.ResponseWriter, r *http.Request) error {
	order, err := dto.NewOrder("dummy customer id for order")
	if err != nil {
		return errtypes.BadRequest{err.Error()}
	}
	log.Println(order.JsonString())
	resp, err := con.SetEntry("AOS_Invoices", nil)
	if resp == nil {
		return nil
	}
	return nil
}
