package handler

import (
	"OpenFabControl/database"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

func Reactivate_user (w http.ResponseWriter, r *http.Request) 	{ user_status(w, r, "activated"); }
func Desactivate_user(w http.ResponseWriter, r *http.Request)	{ user_status(w, r, "desactivated"); }

// only exept new_status to be activated or desactivated (undefined behavior else)
func user_status(w http.ResponseWriter, r *http.Request, new_status string) {
	reject_all_methode_exept(r, w, http.MethodPost);

	var payload struct {
		USERID	int `json:"user_id"`
	}

	if extract_payload_data(r, w, &payload) != nil { return; }

	if !validate_payload(payload.USERID == 0, "user_id can't be empty", w) { return }

	// check status != pending
	if Check_user_status(w, payload.USERID) != nil { return }

	// gen new status
	var toggled string
	if new_status == "activated" { toggled = "desactivated" } else { toggled = "activated" }

	// update table
	query := `UPDATE users SET status = $1 WHERE id = $2`
	_, err := database.Self.Exec(query, toggled, payload.USERID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"msg": "user successfully " + new_status})

}

func Check_user_status(w http.ResponseWriter, id int) error {
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
		if status_check == "pending" {
			http.Error(w, "Cannot desactivate/force-activate an account pending user activation", http.StatusBadRequest)
			return errors.New("pending")
		}
	return nil
}
