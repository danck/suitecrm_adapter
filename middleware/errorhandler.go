package middleware

import (
	"github.com/danck/hawai-suitecrm/errtypes"
	"log"
	"net/http"
)

var (
	// Name of the request's correlation id field
	corrIdKey string = "CorrelationID"
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
			http.Error(w, "whoopsie", http.StatusInternalServerError)
		}

		corrId := r.Header.Get(corrIdKey)
		log.Printf("%s\t%s", corrId, err)
	}
}
