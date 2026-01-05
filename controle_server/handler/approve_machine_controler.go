package handler

import (
	"OpenFabControl/database"
	"fmt"
	"net/http"
)

func Aprove_machine_controler(w http.ResponseWriter, r *http.Request) {
	if reject_all_methode_exept(r, w, http.MethodPost) != nil {
		return
	}

	var payload struct {
		UUID     string `json:"uuid"`
	}

	if extract_payload_data(r, w, &payload) != nil {
		return
	}

	// validate payload data
	if payload.UUID == "" {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	query := "UPDATE machine_controller SET approved = TRUE WHERE uuid = $1"
	result, err := database.Self.Exec(query, payload.UUID)
	if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
	}
	if rows_affected, _ := result.RowsAffected(); rows_affected == 0 {
		http.Error(w, "No device waiting approving with this UUID", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Machine controler approved successfully")
}
