package control

import (
	"invest/service"
	"invest/utils/constants"
	"invest/utils/errormsg"
	"invest/utils/message"
	"net/http"
)

var Get_all_assigned_users_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_all_assigned_users_to_project"

	// parse headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// check whether this is an admin
	if is.RoleName != constants.RoleAdmin {
		message.Respond(w ,r, message.Msg{errormsg.ErrorMethodNotAllowed, 405, fname + " role", "your role name is " + is.RoleName})
	}

	var project_id = service.Get_query_parameter_uint64(r, "project_id", 0)

	// get users
	msg := is.Get_project_with_its_users(project_id)
	msg.Fname = fname + " get"

	message.Respond(w, r, msg)
}
