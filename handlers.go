package main

import (
	"fmt"
	"net/http"
)

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "sdsdf")
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {

}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusNotFound)
}
