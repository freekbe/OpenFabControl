package role_handler

import (
	"OpenFabControl/database"
	"OpenFabControl/utils"
	"net/http"
)

// route to create a role (ex: admin, woodspace, weekend_user, ...)
func Create_role(w http.ResponseWriter, r *http.Request) {

	if utils.Reject_all_methode_exept(r, w, http.MethodPost) != nil { return }

	var payload struct {
		ROLE_NAME	string	`json:"role_name"`
	}

	if utils.Extract_payload_data(r, w, &payload) != nil { return }

	if !utils.Validate_payload(payload.ROLE_NAME == "", "role_name cannot be empty", w) { return }

	// check if role exists
	query := `SELECT EXISTS(SELECT 1 FROM roles WHERE name = $1)`
	var exists bool
	err := database.Self.QueryRow(query, payload.ROLE_NAME).Scan(&exists)
	if err != nil {
		utils.Respond_error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if exists {
		utils.Respond_error(w, "role already exists", http.StatusConflict)
		return
	}

	// create the role in the db
	query = `INSERT INTO roles (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`
	_, err = database.Self.Exec(query, payload.ROLE_NAME)
	if err != nil {
		utils.Respond_error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.Respond_json(w, map[string]any{
		"msg" : "role successfully created",
	}, http.StatusOK)
}
