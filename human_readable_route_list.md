# API access form the machine controler
POST    /machine-api/register                                     // Register a machine controler to the backend
POST    /machine-api/create_user                                  // Create a user

# API access from admin dashboard
## machine controlers
GET     /web-admin-api/get_resource_list_to_approve      // Get the list of machine controlers that requested to be part of the network
GET     /web-admin-api/get_resource_list_approved        // Get the list of machine controlers that are part of the network
POST    /web-admin-api/approve_resource                  // Aprove a machine controler
DELETE  /web-admin-api/delete_resource                   // Delete a machine controler
POST    /web-admin-api/edit_resource                     // Edit the infos of a machine controler
## users
POST    /web-admin-api/create_user                                // Create a users (the send mail dont work)
POST    /web-admin-api/activate                                   // activate an account
POST    /web-admin-api/desactivate                                // desactivate an account
DELETE  /web-admin-api/delete_user                                // Delete a users
POST    /web-admin-api/update_user                                // Update a users
GET     /web-admin-api/get-user-list                              // Get the list of all the user registered in the system 
## roles
POST    /web-admin-api/create_role                                // Create a role
DELETE  /web-admin-api/delete_role                                // Delete a role

# API access from user page
## users
POST    /web-user-api/user_one_time_setup                        // Setup a user (username, password, ect), work only once (after creation)
POST    /web-user-api/login                                      // login trought credentials, return a JWT token




# todo
## users
POST    /web-user-api/me                                          // return profile info
POST    /web-user-api/edit_profile                                // edit self
POST    /web-user-api/logout                                      // logout the logged in user
POST    /web-admin-api/login
POST    /web-admin-api/logout                                     // logout the logged in user

## roles
GET     /web-admin-api/list-roles                                 // Get the list of all the roles created in the system 
POST    /web-admin-api/assign_role_to_user                        // Assign a role to a user
POST    /web-admin-api/remove_role_from_user                      // Remove a role from a user

## machine controlers
POST    /web-admin-api/asign_role_to_user                         // Assign a machine controler to a user
POST    /machine-api/request_session_start                        // Request the start of a machine usage session // to check how to secure the machine communication so you cannot start/end a session for anyone
POST    /machine-api/session_end                                  // end the current session
