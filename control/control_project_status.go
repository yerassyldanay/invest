package control

import (
	"invest/model"
	"invest/service"
	"invest/utils/message"
	"net/http"
)

var Project_get_status_of_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Project_get_status_of_project"
	var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	// parse
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// check whether this user has permission
	msg := is.Check_whether_this_user_can_get_access_to_project_info(project_id)
	msg.Fname = fname + " check"

	if msg.ErrMsg != "" {
		message.Respond(w, r, msg)
		return
	}

	// get status
	var project = model.Project{Id: project_id}
	msg = project.Get_project_with_current_status()
	msg.SetFname(fname, "proj")

	message.Respond(w, r, msg)
}
