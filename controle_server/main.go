package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// project scope
	"OpenFabControl/routes"
	"OpenFabControl/database"
)

func main() {
	fmt.Println("---------------------------------------------")
	fmt.Println("OpenFabControl System [control server]")
	fmt.Println("Author : gazhonsepaskwa")
	fmt.Println("Version : 0.1")
	fmt.Println("---------------------------------------------")
	fmt.Println("")

	database.Initdb()
	defer database.Self.Close()

	routes.Setup_routes()
	runHttpServer()
}

// function that run the http server
func runHttpServer() {
	// check for TLS certs
	if _, errCert := os.Stat("cert.pem"); errCert == nil {
		if _, errKey := os.Stat("key.pem"); errKey == nil {
			addr := os.Getenv("SERVER_ADDR")
			log.Printf("ListenAndServeTLS %s", addr)
			if err := http.ListenAndServeTLS(addr, "cert.pem", "key.pem", nil); err != nil {
				log.Fatalf("ListenAndServeTLS failed: %v", err)
			}
			log.Printf("TLS certs not found cannot continue. bye")
			return
		}
	}
	log.Printf("TLS certs not found cannot continue. bye")
	return
}
