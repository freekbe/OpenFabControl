package utils

import (
	"OpenFabControl/database"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

// return an error if the user pointed by [ id ] have one of the status listed in status_to_check_list []string
func Reject_user_status(w http.ResponseWriter, id int, status_to_check_list []string) error {
	var status_check string
	err := database.Self.QueryRow(`SELECT status FROM users where id = $1`, id).Scan(&status_check)
	if err != nil {
		if err == sql.ErrNoRows {
			Respond_error(w, "No user registered whith this id", http.StatusBadRequest)
		} else {
			Respond_error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return err
	}

	for _, status := range status_to_check_list {
		if status_check == status {
			return errors.New("") // The error is managed by the caller
		}
	}
	return nil
}

// return the error as a json in the folowing format : { "err" : [error msg] }
func Respond_error(w http.ResponseWriter, msg string, status_code int) {
	w.WriteHeader(status_code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"error": msg})
}

// return the success as a json, necesary key pair: "msg" : "..."
func Respond_json(w http.ResponseWriter, json_map map[string]any, status_code int) {
	w.WriteHeader(status_code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(json_map)
}
