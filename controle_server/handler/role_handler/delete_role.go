package role_handler

import (
	"OpenFabControl/utils"
	"OpenFabControl/database"
	"net/http"
)

// route to delete a role (ex: admin, woodspace, weekend_user, ...)
func Delete_role(w http.ResponseWriter, r* http.Request) {

	if utils.Reject_all_methode_exept(r, w, http.MethodDelete) != nil { return }

	var payload struct {
		ROLE_NAME	string	`json:"role_name"`
	}

	if utils.Extract_payload_data(r, w, &payload) != nil { return }

	if !utils.Validate_payload(payload.ROLE_NAME == "", "role_name cannot be empty", w) { return }

	// delete the role from the db
	query := `DELETE FROM roles WHERE name = $1`
	res, err := database.Self.Exec(query, payload.ROLE_NAME)
	if err != nil {
		utils.Respond_error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if rows_affected, _ := res.RowsAffected(); rows_affected == 0 {
		utils.Respond_error(w, "No role with this name saved", http.StatusNotFound)
		return
	}
	utils.Respond_json(w, map[string]any{
		"msg" : "Role deleted successfully",
	}, http.StatusOK)

}
