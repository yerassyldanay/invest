package control

import (
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

var Project_get_status_of_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_get_status_of_project"
	var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	// parse
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// check whether this user has permission
	msg := is.Ganta_check_permission_to_read_ganta(project_id)
	msg.Fname = fname + " check"

	if msg.ErrMsg != "" {
		utils.Respond(w, r, msg); return
	}

	// get status
	var project = model.Project{Id: project_id}
	msg = project.Get_project_with_current_status()

	utils.Respond(w, r, msg)
}
