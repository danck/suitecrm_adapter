package main

import (
	"fmt"
	"net/http"
)

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// TODO(danck) business logic
	}

	fmt.Fprint(w, http.StatusNotFound)
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// TODO(danck) business logic
	}
	if r.Method == "PUT" {
		// TODO(danck) business logic
	}

	fmt.Fprint(w, http.StatusNotFound)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusNotFound)
}
