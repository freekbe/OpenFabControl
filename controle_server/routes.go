package main

import (
	"net/http"
	"fmt"
)

func routes() {
	http.HandleFunc("/register-controller", registerHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the control server!")
	})
}
