package handler

import (
	"OpenFabControl/database"
	"OpenFabControl/utils"
	"encoding/json"
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
	if utils.Reject_user_status(w, payload.USERID, []string{"pending"}) != nil {
		http.Error(w, "Cannot desactivate/force-activate an account pending user activation", http.StatusBadRequest)
		return
	}

	// update table
	query := `UPDATE users SET status = $1 WHERE id = $2`
	_, err := database.Self.Exec(query, new_status, payload.USERID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"msg": "user successfully " + new_status})

}
