package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"

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

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		UUID     string `json:"uuid"`
		Approved *bool  `json:"approved,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if payload.UUID == "" {
		http.Error(w, "missing uuid", http.StatusBadRequest)
		return
	}

	// default approved to false if not provided
	approved := false
	if payload.Approved != nil {
		approved = *payload.Approved
	}

	// insert or update
	query := `INSERT INTO machin_controller (uuid, approved) VALUES ($1, $2)
		ON CONFLICT (uuid) DO UPDATE SET approved = EXCLUDED.approved`
	if _, err := db.Exec(query, payload.UUID, approved); err != nil {
		log.Printf("db insert error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"msg": "registration saved", "uuid": payload.UUID, "approved": approved})
}
