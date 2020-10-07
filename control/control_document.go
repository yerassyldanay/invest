package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

var Document_get = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Document_get"

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// query params
	steps := r.URL.Query()["step"]
	project_id := service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	// security
	msg := is.Check_whether_this_user_can_get_access_to_project_info(project_id)
	msg.Fname = fname + " 1"

	if msg.IsThereAnError() {
		// kick out as a user has nothing to do with the project
		utils.Respond(w, r, msg)
		return
	}

	// logic
	msg = is.Document_get_by_project_id(project_id, steps)
	msg.Fname = fname + " 2"

	utils.Respond(w, r, msg)
}

var Document_add_box_to_upload_document = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Document_add_box"

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// parse request body
	var document = model.Document{}
	if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, " decode")
		return
	}

	// security check
	msg := is.Check_whether_this_user_can_get_access_to_project_info(document.ProjectId)
	if msg.IsThereAnError() {
		utils.Respond(w, r, msg)
		return
	}

	if is.RoleName != utils.RoleExpert && is.RoleName != utils.RoleManager {
		OnlyReturnMethodNotAllowed(w, r, "only spk is allowed. your role: " + is.RoleName, fname, "role")
		return
	}

	msg = is.Add_box_to_upload_document(document)
	msg.SetFname(fname, " add_box")

	utils.Respond(w, r, msg)
}


