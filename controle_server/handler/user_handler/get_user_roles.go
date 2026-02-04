package user_handler

import (
	"OpenFabControl/database"
	"OpenFabControl/model"
	"OpenFabControl/utils"
	"encoding/json"
	"log"
	"net/http"
)

func Get_user_roles(w http.ResponseWriter, r *http.Request) {

	if utils.Reject_all_methode_exept(r, w, http.MethodPost) != nil {
		return
	}

	var payload struct {
		USER_ID int `json:"user_id"`
	}
	payload.USER_ID = -1

	if utils.Extract_payload_data(r, w, &payload) != nil {
		return
	}

	if !utils.Validate_payload(payload.USER_ID == -1, "user_id cannot be empty", w) {
		return
	}

	query := "SELECT r.id, r.name, r.created_at  FROM roles r JOIN users_roles ur ON r.id = ur.role_id WHERE ur.user_id = $1"
	var roles []model.Role
	rows, err := database.Self.Query(query, payload.USER_ID)
	if err != nil {
		utils.Respond_error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// translate the rows
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.ID, &role.NAME, &role.CreatedAt); err != nil {
			utils.Respond_error(w, "internal server error", http.StatusInternalServerError)
			log.Print(err)
			return
		}
		roles = append(roles, role)
	}

	// send data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(roles); err != nil {
		utils.Respond_error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
