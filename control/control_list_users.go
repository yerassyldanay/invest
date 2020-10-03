package control

import (
	"invest/service"
	"invest/utils"
	"net/http"
)

var Get_all_assigned_users_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_all_assigned_users_to_project"

	// parse headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// check whether this is an admin
	if is.RoleName != utils.RoleAdmin {
		utils.Respond(w ,r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " role", "your role name is " + is.RoleName})
	}

	var project_id = service.Get_query_parameter_uint64(r, "project_id", 0)

	// get users
	msg := is.Get_project_with_its_users(project_id)
	msg.Fname = fname + " get"

	utils.Respond(w, r, msg)
}
