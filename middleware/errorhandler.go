package middleware

import (
	"github.com/danck/hawai-suitecrm/errtypes"
	"log"
	"net/http"
)

func ErrorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case errtypes.BadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errtypes.NotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "whoopsie", http.StatusInternalServerError)
		}
	}
}
