package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

/*
	each ganta step has its 'is_done' field, which is an indication of whether
		this step is passed
	this function helps users pass the step manually
 */
var Ganta_confirm_the_ganta_step = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_shift_to_the_next_ganta_step"

	var reqBody = struct {
		ProjectId				uint64				`json:"project_id"`
		Status					string				`json:"status"`
	}{}
	//var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)
	//var status = service.OnlyGetQueryParameter(r, "status", "").(string)
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " json", err.Error()})
		return
	}
	defer r.Body.Close()

	is := service.InvestService{}
	is.OnlyParseRequest(r)

	msg := is.Ganta_change_the_status_of_project(reqBody.ProjectId, reqBody.Status)
	msg.Fname = fname + " status"

	utils.Respond(w, r, msg)
}

var Ganta_change_ganta_time = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_change_ganta_time"

	var ganta = model.Ganta{}
	if err := json.NewDecoder(r.Body).Decode(&ganta); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " json", err.Error()})
		return
	}
	defer r.Body.Close()

	is := service.InvestService{}
	is.OnlyParseRequest(r)

	var msg = model.ReturnMethodNotAllowed("role name is " + is.RoleName)
	if is.RoleName == utils.RoleAdmin {
		msg = ganta.Change_time()
	}

	msg.Fname = fname + " resp"
	utils.Respond(w, r, msg)
}

var Ganta_can_user_change_current_status = func (w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_can_user_change_current_status"

	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	_, msg := is.Ganta_can_user_change_current_status(project_id)
	msg.Fname = fname + " ganta"

	utils.Respond(w, r, msg)
}
