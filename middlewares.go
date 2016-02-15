package main

import (
	"log"
	"net/http"
)

func Tracer(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before")
		h(w, r)
		log.Println("after")
	}
}
