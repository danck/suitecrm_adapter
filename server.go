//TODO(danck) hawai license blabla

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	listenAddress = flag.String("listen-addr", ":8080", "Address to listen on")
	scrmAddr      = flag.String("suitecrm-addr", "http://192.168.29.131/service/v4_1/rest.php", "Address of the SuiteCRM REST api")
	scrmUsr       = flag.String("suitecrm-user", "admin", "Username of SuiteCRM user")
	//scrmPwd       = flag.String("suitecrm-pwd", "21232f297a57a5a743894a0e4a801fc3", "MD5 hash of SuiteCRM password")
	scrmPwd = flag.String("suitecrm-pwd", "admin", "Password of SuiteCRM user")
)

func main() {
	flag.Parse()
	if *scrmAddr == "" {
		fmt.Println("Address for SuiteCRM Endpoint must be set\n Try -h for help")
		return
	}
	con, err := CreateConnection(*scrmAddr, *scrmUsr, *scrmPwd)
	if err != nil {
		fmt.Printf("Connection failed: %s", err)
		return
	}
	log.Printf("SuiteCRM connection established. ID %s", con.SessionId)

	router := http.NewServeMux()
	router.HandleFunc("/customers", Tracer(CustomersHandler))
	router.HandleFunc("/orders", Tracer(OrdersHandler))
	router.HandleFunc("/", Tracer(DefaultHandler))

	log.Printf("Starting to listen on %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, router))
}
