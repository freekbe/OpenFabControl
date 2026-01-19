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
	log.Printf("Starting HTTP server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
