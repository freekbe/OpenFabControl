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
	http.HandleFunc("/machine-api/register", handler.Register)

	// admin page routes
	http.HandleFunc("/web-admin-api/get_machine_controler_list_to_approve",		handler.Get_machine_controler_list_to_approve)
	http.HandleFunc("/web-admin-api/get_machine_controler_list_approved",		handler.Get_machine_controler_list_approved)
	http.HandleFunc("/web-admin-api/approve_machine_controler",					handler.Approve_machine_controler)
	http.HandleFunc("/web-admin-api/delete_machine_controler",					handler.Delete_machine_controler)
	http.HandleFunc("/web-admin-api/edit_machine_controler",					handler.Edit_machine_controler)
	// other
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "controle server working")
	})
}
