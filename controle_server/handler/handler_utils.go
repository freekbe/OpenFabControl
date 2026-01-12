package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
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

// Hash password using bcrypt
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// Compare password with hash
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func validate_payload(condition bool, error_msg string, w http.ResponseWriter) bool {
	if condition {
		http.Error(w, "invalid payload: " + error_msg, http.StatusBadRequest);
		return false
	}
	return true
}
