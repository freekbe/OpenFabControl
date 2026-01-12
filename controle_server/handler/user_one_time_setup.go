package handler

import (
	"OpenFabControl/database"
	"encoding/json"
	"log"
	"net/http"
)

func User_one_time_setup(w http.ResponseWriter, r* http.Request) {

	if reject_all_methode_exept(r, w, http.MethodPost) != nil { return }

	var payload struct {
		VERIF_CODE		string `json:"verification_code"`
		FIRST_NAME 		string `json:"first_name"`
		LAST_NAME		string `json:"last_name"`
		PASSWORD		string `json:"password"`
		TVA				string `json:"tva"`
		ADDRESS			string `json:"facturation_address"`
		ACCOUNT			string `json:"account"`
	}

	if extract_payload_data(r, w, &payload) != nil { return }

	// payload checks
	if payload.VERIF_CODE == "" 	{ http.Error(w, "invalid payload: verification_code cannot be empty", 	http.StatusBadRequest); return }
	if payload.FIRST_NAME == "" 	{ http.Error(w, "invalid payload: first_name cannot be empty", 			http.StatusBadRequest); return }
	if payload.LAST_NAME == "" 		{ http.Error(w, "invalid payload: last_name cannot be empty",			http.StatusBadRequest); return }
	if payload.PASSWORD == "" 		{ http.Error(w, "invalid payload: password cannot be empty", 			http.StatusBadRequest); return }
	if payload.ACCOUNT == "" 		{ http.Error(w, "invalid payload: account cannot be empty", 			http.StatusBadRequest); return }

	// check if the validation code exist and that the account is waiting setup
	var status string
	err := database.Self.QueryRow(`SELECT status FROM users WHERE verification_code = $1`, payload.VERIF_CODE).Scan(&status)
	if err != nil {
		http.Error(w, "Invalid verification code", http.StatusBadRequest)
		return
	}
	if status != "pending" {
		http.Error(w, "Account already set-up", http.StatusBadRequest)
		return
	}

	// set information in db
	hashed_password, err := hashPassword(payload.PASSWORD);
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	query := `UPDATE users SET password = $1, first_name = $2, last_name = $3, tva = $4, facturation_address = $5, account = $6, status = $7 WHERE verification_code = $8`
	_, err = database.Self.Exec(query, hashed_password, payload.FIRST_NAME, payload.LAST_NAME, payload.TVA, payload.ADDRESS, payload.ACCOUNT, "active", payload.VERIF_CODE)
	if err != nil {
		log.Printf("db update error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"msg": "user set-up, you can login now"})
}
