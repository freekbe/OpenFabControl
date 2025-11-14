package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	fmt.Println("---------------------------------------------")
	fmt.Println("OpenFabControl System [control server]")
	fmt.Println("Author : gazhonsepaskwa")
	fmt.Println("Version : 0.1")
	fmt.Println("---------------------------------------------")
	fmt.Println("")

	initdb()
	defer db.Close()

	routes()
	runHttpServer()

	return
}

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
