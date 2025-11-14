package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

// all the routes
func routes_handler() {
	http.HandleFunc("/register-controller", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("Received a POST request")

			// read JSON data
			var data map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			// Process the data
			if value, ok := data["uuidv4"]; ok {
				fmt.Println("Received data:", value)
				w.Header().Set("Content-Type", "application/json")
				response := map[string]string{
					"msg": "registration accepted",
				}
				json.NewEncoder(w).Encode(response)
			} else {
				response := map[string]string{
					"msg": "registration rejected",
				}
				json.NewEncoder(w).Encode(response)
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the control server!")
	})
}

func main() {
	// information
	fmt.Println("Starting control server...")
	fmt.Println("Author : gazhonsepaskwa")
	fmt.Println("Version : 0.1")

	routes_handler();

	// start server
	err := http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
	if err != nil { log.Fatal("ListenAndServeTLS failed:", err) }
}
