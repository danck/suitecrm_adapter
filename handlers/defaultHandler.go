package handlers

import (
	"github.com/danck/hawai-suitecrm/errtypes"
	"net/http"
)

// DefaultHandler gets called to handle requests that don't match a route
func DefaultHandler(w http.ResponseWriter, r *http.Request) error {
	return errtypes.NotImplemented{"No matching prefix"}
}
