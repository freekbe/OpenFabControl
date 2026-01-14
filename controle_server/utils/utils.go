package utils

import (
	"OpenFabControl/database"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
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

// helper that hash password using bcrypt
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// helper that compare password with hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
