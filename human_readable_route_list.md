# API access form the machine controler
POST /machine-api/register                                     // Register a machine controler to the backend

# API access from admin dashboard
GET /web-admin-api/get_machine_controler_list_to_approve      // Get the list of machine controlers that requested to be part of the network
GET /web-admin-api/get_machine_controler_list_approved        // Get the list of machine controlers that are part of the network
GET /web-admin-api/approve_machine_controler                  // Aprove a machine controler
GET /web-admin-api/delete_machine_controler                   // Delete a machine controler

# todo
POST /web-admin-api/edit_machine_controler                     // Edit the infos of a machine controler
