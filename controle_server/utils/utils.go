package utils

import (
	"OpenFabControl/database"
	"net/http"
	"database/sql"
	"errors"
)

func Reject_user_status(w http.ResponseWriter, id int, status_to_check_list []string) error {
	var status_check string
	err := database.Self.QueryRow(`SELECT status FROM users where id = $1`, id).Scan(&status_check)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No user registered whith this id", http.StatusBadRequest)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
