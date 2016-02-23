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

		var errorName string

		switch err.(type) {
		case errtypes.BadRequest:
			errorName = "BadRequest"
			http.Error(w, errorName+err.Error(), http.StatusBadRequest)
		case errtypes.NotFound:
			errorName = "NotFound"
			http.Error(w, errorName+err.Error(), http.StatusNotFound)
		case errtypes.NotImplemented:
			errorName = "NotImplemented"
			http.Error(w, errorName+err.Error(), http.StatusNotImplemented)
		default:
			errorName = "InternalError"
			http.Error(w, "whoopsie", http.StatusInternalServerError)
		}

		corrId := r.Header.Get(corrIdKey)
		log.Printf("%s\t%s\t%s", corrId, errorName, err)
	}
}
