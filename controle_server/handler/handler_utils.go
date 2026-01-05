package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func reject_all_methode_exept(r *http.Request, w http.ResponseWriter, methode string) error {
	if r.Method != methode {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

func extract_payload_data(r *http.Request, w http.ResponseWriter, payload any) error {
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return fmt.Errorf("invalid json")
	}
	return nil
}
