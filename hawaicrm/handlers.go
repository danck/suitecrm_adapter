package hawaicrm

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// CustomersHandler is the exposed end point for the customers prefix
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

// OrdersHandler is the exposed endpoint for the orders prefix
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

// defaultHandler gets calles if no other handler matches
func defaultHandler(w http.ResponseWriter, r *http.Request) error {
	return NotImplemented{"Prefix doesn't exist"}
}

// markAsPaid requires an order/invoice id and sets its status to 'Paid'
func markAsPaid(w http.ResponseWriter, r *http.Request) error {
	var o Order
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&o)
	if err != nil || o.ID == "" {
		return BadRequest{err.Error()}
	}

	orderData := []KeyValuePair{
		KeyValuePair{"id", o.ID},
		KeyValuePair{"status", "Paid"},
	}

	_, err = crmSetEntry("AOS_Invoices", orderData)
	if err != nil {
		return err
	}

	w.WriteHeader(200)

	return nil
}

// createOrder creates a new order with the given data and associates it with
// an existing customer
func createOrder(w http.ResponseWriter, r *http.Request) error {
	type Message struct {
		Order    Order
		Customer Customer
	}
	var msg Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		return BadRequest{err.Error()}
	}

	totalAmount := strconv.FormatFloat(msg.Order.Price, 'f', -1, 64)
	shippingAmount := strconv.FormatFloat(msg.Order.ShippingCosts, 'f', -1, 64)

	orderData := []KeyValuePair{
		KeyValuePair{"total_amount", totalAmount},
		KeyValuePair{"billing_account_id", msg.Customer.ID},
		KeyValuePair{"shipping_amount", shippingAmount},
		KeyValuePair{"name", "GENERATED BY API"},
		KeyValuePair{"status", "Unpaid"},
	}

	orderResponse, err := crmSetEntry("AOS_Invoices", orderData)
	if err != nil {
		return err
	}

	var orderMap map[string]interface{}
	err = json.Unmarshal(orderResponse, &orderMap)
	if err != nil {
		return err
	}

	orderID, ok := orderMap["id"]
	if !ok {
		return errors.New("Unable to acquire order id")
	}

	msg.Order.ID = orderID.(string)
	json.NewEncoder(w).Encode(msg.Order)

	return nil
}

func createCustomer(w http.ResponseWriter, r *http.Request) error {
	var c Customer

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	if err != nil {
		return BadRequest{err.Error()}
	}

	// Mapping between HAWAI 'customer' and SuiteCRM 'account' type
	customerData := []KeyValuePair{
		KeyValuePair{"name", c.FirstName + " " + c.LastName},
		KeyValuePair{"email1", c.Email},
		// Password is being ignored for now
	}

	// Mapping between HAWAI 'address' and SuiteCRM 'contact' type
	addressData := []KeyValuePair{
		KeyValuePair{"description", "Primary Contact"},
		KeyValuePair{"first_name", c.Address.FirstName},
		KeyValuePair{"last_name", c.Address.LastName},
		KeyValuePair{"primary_address_street", c.Address.Street + " " + c.Address.HouseNumber},
		KeyValuePair{"primary_address_city", c.Address.City},
		KeyValuePair{"primary_address_postalcode", c.Address.Zipcode},
		KeyValuePair{"primary_address_country", c.Address.Country},
	}

	// Mapping between HAWAI 'address' and SuiteCRM 'contact' type
	shipmentAddressData := []KeyValuePair{
		KeyValuePair{"description", "Shipment Address"},
		KeyValuePair{"first_name", c.ShipmentAddress.FirstName},
		KeyValuePair{"last_name", c.ShipmentAddress.LastName},
		KeyValuePair{"primary_address_street", c.ShipmentAddress.Street + " " + c.Address.HouseNumber},
		KeyValuePair{"primary_address_city", c.ShipmentAddress.City},
		KeyValuePair{"primary_address_postalcode", c.ShipmentAddress.Zipcode},
		KeyValuePair{"primary_address_country", c.ShipmentAddress.Country},
	}

	// Mapping between HAWAI 'address' and SuiteCRM 'contact' type
	invoiceAddressData := []KeyValuePair{
		KeyValuePair{"description", "Invoice Address"},
		KeyValuePair{"first_name", c.InvoiceAddress.FirstName},
		KeyValuePair{"last_name", c.InvoiceAddress.LastName},
		KeyValuePair{"primary_address_street", c.InvoiceAddress.Street + " " + c.Address.HouseNumber},
		KeyValuePair{"primary_address_city", c.InvoiceAddress.City},
		KeyValuePair{"primary_address_postalcode", c.InvoiceAddress.Zipcode},
		KeyValuePair{"primary_address_country", c.InvoiceAddress.Country},
	}

	customerResponse, err := crmSetEntry("Accounts", customerData)
	if err != nil {
		return err
	}
	addressResponse, err := crmSetEntry("Contacts", addressData)
	if err != nil {
		return err
	}
	shipmentResponse, err := crmSetEntry("Contacts", shipmentAddressData)
	if err != nil {
		return err
	}
	invoiceResponse, err := crmSetEntry("Contacts", invoiceAddressData)
	if err != nil {
		return err
	}

	var customerMap map[string]interface{}
	var addressMap map[string]interface{}
	var shipmentMap map[string]interface{}
	var invoiceMap map[string]interface{}

	err = json.Unmarshal(customerResponse, &customerMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(addressResponse, &addressMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(shipmentResponse, &shipmentMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(invoiceResponse, &invoiceMap)
	if err != nil {
		return err
	}

	customerID, ok := customerMap["id"]
	if !ok {
		return errors.New("Not able to retrieve customer ID")
	}
	addressID, ok := addressMap["id"]
	if !ok {
		return errors.New("Not able to retrieve address ID")
	}
	shipmentID, ok := shipmentMap["id"]
	if !ok {
		return errors.New("Not able to retrieve shipment ID")
	}
	invoiceID, ok := invoiceMap["id"]
	if !ok {
		return errors.New("Not able to retrieve invoice ID")
	}

	relatedIDs := []string{
		addressID.(string),
		shipmentID.(string),
		invoiceID.(string)}
	values := []KeyValuePair{}

	relationResponse, err := crmSetRelationship(
		"Accounts",
		customerID.(string),
		"contacts",
		relatedIDs,
		values)

	var relationResponseMap map[string]int
	err = json.Unmarshal(relationResponse, &relationResponseMap)
	if err != nil || relationResponseMap["failed"] != 0 {
		return errors.New("Adding relationship failed: " + customerID.(string))
	}

	c.ID = customerID.(string)
	json.NewEncoder(w).Encode(c)
	return nil
}
