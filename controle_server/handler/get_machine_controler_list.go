package handler

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"log"

	"OpenFabControl/database"
	"OpenFabControl/model"
)

func Get_machine_controler_list_approved(w http.ResponseWriter, r *http.Request)	{ get_machine_controler(w,r,true)  }
func Get_machine_controler_list_to_approve(w http.ResponseWriter, r *http.Request)	{ get_machine_controler(w,r,false) }

func get_machine_controler(w http.ResponseWriter, r *http.Request, approved bool) {
	reject_all_methode_exept(r, w, http.MethodGet)

	// get the controllers
	query := ""
	if approved	{ query = "SELECT * FROM machine_controller WHERE approved = TRUE"
	} else		{ query = "SELECT * FROM machine_controller WHERE approved = FALSE" }
	var controllers []model.Machine_controller
	rows, err := database.Self.Query(query);
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound);
			return
		}
	}
	defer rows.Close()

	// translate the rows
	for rows.Next() {
		var controller model.Machine_controller
		if err := rows.Scan(&controller.ID,
							&controller.UUID,
							&controller.ZONE,
							&controller.NAME,
							&controller.MANUAL,
							&controller.PRICE_BOOKING_IN_EUR,
							&controller.PRICE_USAGE_IN_EUR,
							&controller.Approved,
							&controller.CreatedAt); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		controllers = append(controllers, controller)
	}

	// headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content_Type", "application/json")

	// send data
	if err := json.NewEncoder(w).Encode(controllers); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
