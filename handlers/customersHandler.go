package handlers

import (
	"encoding/json"
	"fmt"
	con "gitlab.com/danck/hawai-suitecrm/connector"
	"gitlab.com/danck/hawai-suitecrm/dto"
	"gitlab.com/danck/hawai-suitecrm/errtypes"
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

	return errtypes.NotImplemented{"Prefix exists, but no matching HTTP method"}
}

func createCustomer(w http.ResponseWriter, r *http.Request) error {
	log.Println("Creating customer")
	var c dto.Customer
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	if err != nil {
		log.Println("Decode failed", err)
		return errtypes.BadRequest{err.Error()}
	}

	//log.Println("Decoded request body: ", c.String())
	nameValueList := []con.KeyValuePair{
		con.KeyValuePair{"name", c.FirstName + " " + c.LastName},
	}

	resp, err := con.SetEntry("Accounts", nameValueList)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, resp.(string))
	//log.Println("customer:", customer.JsonString())
	//fmt.Fprintf(w, customer.JsonString())
	return nil
}
