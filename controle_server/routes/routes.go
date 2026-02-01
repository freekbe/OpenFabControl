package routes

import (
	"OpenFabControl/handler/machine_controler_handler"
	"OpenFabControl/handler/role_handler"
	"OpenFabControl/handler/user_handler"
	"fmt"
	"net/http"
)

// function to manage routes
func Setup_routes() {
	//////////////////////////////
	// machine controler routes //
	//////////////////////////////

	http.HandleFunc("/machine-api/register", 									machine_controler_handler.Register)
	http.HandleFunc("/machine-api/create_user", 								user_handler.Create_user)

	///////////////////////
	// admin page routes //
	///////////////////////

	// TODO : admin pages have to be protected (not done for the moment for dev purpose)

	// machine controlers
	http.HandleFunc("/web-admin-api/get_machine_controler_list_to_approve",		machine_controler_handler.Get_machine_controler_list_to_approve)
	http.HandleFunc("/web-admin-api/get_machine_controler_list_approved",		machine_controler_handler.Get_machine_controler_list_approved)
	http.HandleFunc("/web-admin-api/approve_machine_controler",					machine_controler_handler.Approve_machine_controler)
	http.HandleFunc("/web-admin-api/delete_machine_controler",					machine_controler_handler.Delete_machine_controler)
	http.HandleFunc("/web-admin-api/edit_machine_controler",					machine_controler_handler.Edit_machine_controler)

	// users
	http.HandleFunc("/web-admin-api/create_user", 								user_handler.Create_user)
	http.HandleFunc("/web-admin-api/delete_user", 								user_handler.Delete_user)
	http.HandleFunc("/web-admin-api/update_user",	 							user_handler.Update_user)
	// http.HandleFunc("/web-admin-api/get_user_list",	 						user_handler.Get_user_list)
	http.HandleFunc("/web-admin-api/desactivate_user",	 						user_handler.Desactivate_user)
	http.HandleFunc("/web-admin-api/reactivate_user",	 						user_handler.Reactivate_user)
	// roles
	http.HandleFunc("/web-admin-api/create_role", 								role_handler.Create_role)
	http.HandleFunc("/web-admin-api/delete_role", 								role_handler.Delete_role)

	///////////////////////
	// user pages routes //
	///////////////////////

	http.HandleFunc("/web-user-api/user_one_time_setup", 						user_handler.User_one_time_setup)
	http.HandleFunc("/web-user-api/login",				 						user_handler.Login)

	///////////
	// other //
	///////////

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "controle server working")
	})
}
