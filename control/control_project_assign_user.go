package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils/constants"
	"invest/utils/errormsg"
	"invest/utils/message"

	"net/http"
)

var Assign_user_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Admin_assign_user_to_project"
	var pu = model.ProjectsUsers{}

	if err := json.NewDecoder(r.Body).Decode(&pu); err  != nil {
		message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " 1", err.Error()})
		return
	}
	defer r.Body.Close()

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// security check
	if is.RoleName != constants.RoleAdmin {
		message.Respond(w, r, message.Msg{errormsg.ErrorMethodNotAllowed, 405, fname + " role", "not admin"})
		return
	}

	// check project
	if msg := is.Does_project_exist(pu.ProjectId); msg.IsThereAnError() {
		message.Respond(w, r, msg)
		return
	}

	// check user
	if msg := is.Does_user_has_given_role(pu.UserId, []string{constants.RoleExpert, constants.RoleManager}); msg.IsThereAnError() {
		message.Respond(w, r, msg)
		return
	}

	// logic
	msg := is.Assign_user_to_project(pu)
	msg.Fname = fname + " asg"

	message.Respond(w, r, msg)
}

/*
	remove user & project relation
 */
var Remove_user_from_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Remove_user_from_project"
	var pu = model.ProjectsUsers{}

	if err := json.NewDecoder(r.Body).Decode(&pu); err != nil {
		message.Respond(w, r, message.Msg{
			Message: errormsg.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// logic
	msg := is.Assign_remove_relation(pu)
	message.Respond(w, r, msg)
}

