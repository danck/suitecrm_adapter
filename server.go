//TODO(danck) hawai license blabla

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	listenAddress = flag.String("listenAddr", ":8080", "Address to listen on")
	scrmAddr      = flag.String("suitcrmAddr", "", "Address of the SuiteCRM REST api")
)

func main() {
	flag.Parse()
	if *scrmAddr == "" {
		fmt.Println("Address for SuiteCRM Endpoint must be set\n Try -h for help")
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("/customers", Tracer(CustomersHandler))
	router.HandleFunc("/orders", Tracer(OrdersHandler))
	router.HandleFunc("/", Tracer(DefaultHandler))

	log.Fatal(http.ListenAndServe(*listenAddress, router))
}
