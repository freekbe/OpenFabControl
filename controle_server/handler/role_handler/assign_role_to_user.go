package role_handler

import (
	"OpenFabControl/database"
	"OpenFabControl/utils"
	"fmt"
	"net/http"
)

func Assign_role_to_user(w http.ResponseWriter, r *http.Request) {

	if utils.Reject_all_methode_exept(r, w, http.MethodPost) != nil {
		return
	}

	var payload struct {
		USER_ID int `json:"user_id"`
		ROLE_ID int `json:"role_id"`
	}
	payload.USER_ID = -1
	payload.ROLE_ID = -1

	if utils.Extract_payload_data(r, w, &payload) != nil {
		return
	}

	// validate payload
	if !utils.Validate_payload(payload.USER_ID == -1, "user_id cannot be empty", w) {
		return
	}
	if !utils.Validate_payload(payload.ROLE_ID == -1, "role_id cannot be empty", w) {
		return
	}

	// Insert relation in the junction table
	querry := `INSERT INTO users_roles (user_id, role_id) VALUES ($1, $2) ON CONFLICT (user_id, role_id) DO NOTHING`
	res, err := database.Self.Exec(querry, payload.USER_ID, payload.ROLE_ID)
	if err != nil {
		utils.Respond_error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	// check if role is not already assigned to user
	v, err := res.RowsAffected()
	if err != nil {
		utils.Respond_error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if v == 0 {
		utils.Respond_error(w, "This role is already assigned to this user", http.StatusInternalServerError)
		return
	}

	utils.Respond_json(w, map[string]any{
		"msg": fmt.Sprintf("role with id %v successfully added to user with id %v", payload.ROLE_ID, payload.USER_ID),
	}, http.StatusOK)
}
