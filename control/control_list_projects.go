package control

import (
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

/*
	user (manager, lawyer, financier or others) can get ptojects that have been assigned to them
*/
var Get_own_projects = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_project_get_all"

	// parse headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// prepare status & step
	statuses, steps := service.OnlyPrepareStatusAndStep(r)

	// get projects (provide offset)
	msg := is.Get_own_projects(statuses, steps)
	msg.Fname = fname + " own"

	utils.Respond(w, r, msg)
}

/*
	admin can get all projects with users assigned to them
 */
var Get_all_projects_by_user_and_status = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_project_get_all"

	// parse headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// parse query parameters
	var user_id = service.OnlyGetQueryParameter(r, "user_id", uint64(0)).(uint64)

	// convert the external status to the internal ones
	statuses, _ := service.OnlyPrepareStatusAndStep(r)

	// logic
	msg := is.Get_projects_by_user_id_and_status(user_id, statuses, []int{1, 2})
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}

var Get_all_projects_by_statuses = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_all_projects"

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check
	if is.RoleName != utils.RoleAdmin {
		msg := model.ReturnMethodNotAllowed("only admin. your role is " + is.RoleName)
		msg.Fname = fname + " role"
		utils.Respond(w, r, msg)
		return
	}

	// parse parameters
	statuses, steps := service.OnlyPrepareStatusAndStep(r)

	// logic
	msg := is.Get_all_projects_by_statuses(statuses, steps)
	msg.Fname = fname + " get"

	utils.Respond(w, r, msg)
}
