package role_handler

import (
	"OpenFabControl/database"
	"OpenFabControl/model"
	"OpenFabControl/utils"
	"encoding/json"
	"log"
	"net/http"
)

// handler for the get user route
func Get_role_list(w http.ResponseWriter, r *http.Request) {

	if utils.Reject_all_methode_exept(r, w, http.MethodGet) != nil {
		return
	}

	// get the users
	query := "SELECT id, name, created_at FROM roles"
	var roles []model.Role
	rows, err := database.Self.Query(query)
	if err != nil {
		utils.Respond_error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// translate the rows
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.ID,
			&role.NAME,
			&role.CreatedAt); err != nil {
			utils.Respond_error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("%v", err)
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
