package control

import (
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

/*
	which:
		* parents - ganta steps, which are steps of a process
		* children - ganta sub-steps, which are documents (related to one document)
 */
var Ganta_restricted_get_help = func(which string, w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_get_parent_ganta_steps"
	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// check permission
	var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	// security check
	msg := is.Check_whether_this_user_can_get_access_to_project_info(project_id)
	msg.Fname = fname + " is"

	if msg.ErrMsg != "" {
		utils.Respond(w, r, msg)
		return
	}

	// get current status & step
	var project = model.Project{Id: project_id}
	err := project.GetAndUpdateStatusOfProject(model.GetDB())

	if err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInternalDbError, 417, fname + " db", err.Error()})
		return
	}

	var ganta = model.Ganta{ProjectId: project_id}
	switch which {
	case "parents":
		msg = ganta.Get_parent_ganta_steps_by_project_id_and_step(project.Step)
	case "children":
		msg = ganta.Get_child_ganta_steps_by_project_id_and_step(project.Step)
	default:
		msg = model.ReturnMethodNotAllowed("you are requesting " + which)
	}

	msg.Fname = fname + " ganta"
	utils.Respond(w, r, msg)
}

/*
	restricted - because you will get ganta steps either for project step / stage 1 or 2
 */
var Ganta_restricted_get_parent_ganta_steps = func(w http.ResponseWriter, r *http.Request) {
	Ganta_restricted_get_help("parents", w, r)
}

var Ganta_restricted_get_child_ganta_steps = func(w http.ResponseWriter, r *http.Request) {
	Ganta_restricted_get_help("children", w, r)
}


