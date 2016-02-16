//TODO(danck) hawai license blabla

package main

import (
	"flag"
	mw "github.com/danck/hawai-suitecrm/middleware"
	"log"
	"net/http"
)

// Command line parameters
var (
	listenAddress = flag.String(
		"listen-addr",
		":8080",
		"Address to listen on")
	scrmAddr = flag.String(
		"suitecrm-addr",
		"http://localhost/service/v4_1/rest.php",
		"Address of the SuiteCRM REST api")
	scrmUsr = flag.String(
		"suitecrm-user",
		"admin",
		"Username of SuiteCRM user")
	scrmPwd = flag.String(
		"suitecrm-pwd",
		"admin",
		"Password of SuiteCRM user")
)

// Global services
var (
	con *Connection
)

func main() {
	flag.Parse()

	// Connect to SuiteCRM
	con, err := NewConnection(*scrmAddr, *scrmUsr, *scrmPwd)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SuiteCRM connection established. ID %s", con.SessionId)

	// Initialize datatype managers

	// Register handlers
	router := http.NewServeMux()
	router.HandleFunc("/customers", mw.ErrorHandler(CustomersHandler))
	router.HandleFunc("/orders", mw.ErrorHandler(OrdersHandler))
	router.HandleFunc("/", mw.ErrorHandler(DefaultHandler))

	// Start the server
	log.Printf("Starting to listen on %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, router))
}