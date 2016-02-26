//TODO(danck) hawai license blabla

package hawaicrm

import (
	"flag"
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
	logFile = flag.String(
		"log-file",
		"hawai-suitecrm.log",
		"Log file")
)

func Main() {
	flag.Parse()

	// Connect to SuiteCRM
	err := Connect(*scrmAddr, *scrmUsr, *scrmPwd)
	if err != nil {
		log.Fatal(err)
	}

	// Register handlers
	router := http.NewServeMux()
	router.HandleFunc("/customers", ErrorHandler(CustomersHandler))
	router.HandleFunc("/orders", ErrorHandler(OrdersHandler))
	router.HandleFunc("/", ErrorHandler(defaultHandler))

	// Start the server
	log.Printf("Starting to listen on %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, router))
}