package routes

import (
	"fmt"
	"net/http"

	// project scope
	"OpenFabControl/handler"
)

// function to manage routes
func Setup_routes() {
	//////////////////////////////
	// machine controler routes //
	//////////////////////////////

	http.HandleFunc("/machine-api/register", 									handler.Register)
	http.HandleFunc("/machine-api/create_user", 								handler.Create_user)

	///////////////////////
	// admin page routes //
	///////////////////////

	// TODO : admin pages have to be protected (not done for the moment for dev purpose)

	// machine controlers
	http.HandleFunc("/web-admin-api/get_machine_controler_list_to_approve",		auth_middleware(handler.Get_machine_controler_list_to_approve))
	http.HandleFunc("/web-admin-api/get_machine_controler_list_approved",		handler.Get_machine_controler_list_approved)
	http.HandleFunc("/web-admin-api/approve_machine_controler",					handler.Approve_machine_controler)
	http.HandleFunc("/web-admin-api/delete_machine_controler",					handler.Delete_machine_controler)
	http.HandleFunc("/web-admin-api/edit_machine_controler",					handler.Edit_machine_controler)

	// users
	http.HandleFunc("/web-admin-api/create_user", 								handler.Create_user)
	// http.HandleFunc("/web-admin-api/delete_user", 							handler.Delete_user)
	// http.HandleFunc("/web-admin-api/update_user",	 						handler.Update_user)
	// http.HandleFunc("/web-admin-api/get_user_list",	 						handler.Get_user_list)
	http.HandleFunc("/web-admin-api/desactivate_user",	 						handler.Desactivate_user)
	http.HandleFunc("/web-admin-api/reactivate_user",	 						handler.Reactivate_user)
	// roles

	///////////////////////
	// user pages routes //
	///////////////////////

	http.HandleFunc("/web-user-api/user_one_time_setup", 						handler.User_one_time_setup)
	http.HandleFunc("/web-user-api/login",				 						handler.Login)

	///////////
	// other //
	///////////

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "controle server working")
	})
}
