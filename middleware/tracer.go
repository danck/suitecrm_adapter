package middleware

import (
	"log"
	"net/http"
)

func tracer(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before")
		h(w, r)
		log.Println("after")
	}
}
