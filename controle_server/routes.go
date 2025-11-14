package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"database/sql"
)

func routes() {
	http.HandleFunc("/register-controller", registerHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the control server!")
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// reject non POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		UUID     string `json:"uuid"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// validate payload
	if payload.UUID == "" {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}


	query := `INSERT INTO machin_controller (uuid, approved) VALUES ($1, $2)
	ON CONFLICT (uuid) DO NOTHING`

	// Check if UUID already exists
	var existingUUID string
	err := db.QueryRow(`SELECT uuid FROM machin_controller WHERE uuid = $1`, payload.UUID).Scan(&existingUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			// UUID doesn't exist, proceed with insertion
			_, err := db.Exec(query, payload.UUID, false)
			if err != nil {
				log.Printf("db insert error: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{"msg": "registration saved", "uuid": payload.UUID})
			return
		} else {
			// Other database error
			log.Printf("db query error: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	} else {
		// UUID already exists
		http.Error(w, "UUID already registered", http.StatusBadRequest)
		return
	}
}
