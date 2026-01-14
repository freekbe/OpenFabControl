package handler

import (
	"OpenFabControl/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// helper to reject a request if the methode used is not the specified one
func reject_all_methode_exept(r *http.Request, w http.ResponseWriter, methode string) error {
	if r.Method != methode {
		utils.Respond_error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

// helper to extract the json data sent
func extract_payload_data(r *http.Request, w http.ResponseWriter, payload any) error {
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		utils.Respond_error(w, "invalid json", http.StatusBadRequest)
		return fmt.Errorf("invalid json")
	}
	return nil
}

// helper that hash password using bcrypt
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// helper that compare password with hash
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// helper to check a condition of the payload
func validate_payload(condition bool, error_msg string, w http.ResponseWriter) bool {
	if condition {
		utils.Respond_error(w, "invalid payload: " + error_msg, http.StatusBadRequest)
		return false
	}
	return true
}
