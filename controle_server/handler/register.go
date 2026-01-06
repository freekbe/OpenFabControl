package handler

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"log"

	"OpenFabControl/database"
)

// register a new machine controler
func Register(w http.ResponseWriter, r *http.Request) {
	reject_all_methode_exept(r, w, http.MethodPost)

	var payload struct {
		UUID    string `json:"uuid"`
		NAME	string `json:"name"`
	}

	// extract payload data
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// validate payload data
	if payload.UUID == "" || payload.NAME == "" {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO machine_controller (uuid, zone, name, manual, price_booking_in_eur, price_usage_in_eur, approved) VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (uuid) DO NOTHING`

	// Check if UUID already exists
	var existingUUID string
	err := database.Self.QueryRow(`SELECT uuid FROM machine_controller WHERE uuid = $1`, payload.UUID).Scan(&existingUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			// UUID doesn't exist, proceed with insertion
			_, err := database.Self.Exec(query, payload.UUID, "UNDEFINED", payload.NAME, "UNDEFINED", 0, 0, false)
			if err != nil {
				log.Printf("db insert error: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{"msg": "registration saved", "uuid": payload.UUID, "name": payload.NAME})
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
