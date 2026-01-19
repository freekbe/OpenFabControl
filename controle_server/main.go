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
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":3000"
	}
	
	// Try TLS first if certs are available, otherwise fall back to HTTP
	if _, errCert := os.Stat("cert.pem"); errCert == nil {
		if _, errKey := os.Stat("key.pem"); errKey == nil {
			log.Printf("TLS certs found, starting HTTPS server on %s", addr)
			if err := http.ListenAndServeTLS(addr, "cert.pem", "key.pem", nil); err != nil {
				log.Fatalf("ListenAndServeTLS failed: %v", err)
			}
			return
		}
	}
	
	// Fall back to HTTP (for use behind reverse proxy with SSL termination)
	log.Printf("TLS certs not found, starting HTTP server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
