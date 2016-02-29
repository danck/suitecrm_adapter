package hawaicrm

import (
	"log"
	"net/http"
)

var (
	// Name of the request's correlation id field
	corrIDKey = "CorrelationID"
)

// ErrorHandler is a wrapper for types of HandlerFunc that logs returned errors
// and reports them to the requestor
func ErrorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		corrID := r.Header.Get(corrIDKey)
		log.Printf("Receiving request %s", corrID)

		err := f(w, r)

		if err != nil {
			var errorType string

			// Publishes the error message (if any) to the requestor
			switch err.(type) {
			case BadRequest:
				errorType = "BadRequest"
				http.Error(w, errorType+": "+err.Error(), http.StatusBadRequest)
			case NotFound:
				errorType = "NotFound"
				http.Error(w, errorType+": "+err.Error(), http.StatusNotFound)
			case NotImplemented:
				errorType = "NotImplemented"
				http.Error(w, errorType+": "+err.Error(), http.StatusNotImplemented)
			default:
				errorType = "InternalError"
				http.Error(w, "whoopsie", http.StatusInternalServerError)
			}

			// Log the error message (if any) internally
			log.Printf("%s\t%s\t%s", corrID, errorType, err)
		}
	}
}
