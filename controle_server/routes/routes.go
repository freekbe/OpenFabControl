package routes

import (
	"fmt"
	"net/http"

	// project scope
	"OpenFabControl/handler"
)

// function to manage routes
func Setup_routes() {
	// machine controler routes
	http.HandleFunc("/machine-api/machine_controler/register", handler.Register)

	// admin page routes
	http.HandleFunc("/web-api/ofc_admin/get_machine_controler_list_to_approve",		handler.Get_machine_controler_list_to_approve)
	http.HandleFunc("/web-api/ofc_admin/get_machine_controler_list_approved",		handler.Get_machine_controler_list_approved)
	http.HandleFunc("/web-api/ofc_admin/approve_machine_controler",					handler.Approve_machine_controler)
	http.HandleFunc("/web-api/ofc_admin/delete_machine_controler",					handler.Delete_machine_controler)
	// other
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "controle server working")
	})
}
