package app

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

/*
	which:
		* parents - ganta steps, which are steps of a process
		* children - ganta sub-steps, which are documents (related to one document)
*/
var GantaRestrictedGetHelp = func(which string, w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_get_parent_ganta_steps"
	var is = service.InvestService{}
	is.OnlyParseRequest(r)

	// check permission
	var project_id = service.OnlyGetQueryParameter(r, "project_id", uint64(0)).(uint64)

	// security check
	msg := is.CheckWhetherThisUserCanGetAccessToProjectInfo(project_id)
	msg.Fname = fname + " is"

	if msg.ErrMsg != "" {
		message.Respond(w, r, msg)
		return
	}

	// get current status & step
	var project = model.Project{Id: project_id}
	err := project.GetAndUpdateStatusOfProject(model.GetDB())

	if err != nil {
		message.Respond(w, r, message.Msg{errormsg.ErrorInternalDbError, 417, fname + " database", err.Error()})
		return
	}

	var ganta = model.Ganta{ProjectId: project_id}
	switch which {
	case "parents":
		msg = ganta.GetParentGantaStepsByProjectIdAndStep(project.Step)
	case "children":
		msg = ganta.GetChildGantaStepsByProjectIdAndStep(project.Step)
	default:
		msg = model.ReturnMethodNotAllowed("you are requesting " + which)
	}

	msg.Fname = fname + " ganta"
	message.Respond(w, r, msg)
}

/*
	restricted - because you will get ganta steps either for project step / stage 1 or 2
*/
var GantaRestrictedGetParentGantaSteps = func(w http.ResponseWriter, r *http.Request) {
	GantaRestrictedGetHelp("parents", w, r)
}

var GantaRestrictedGetChildGantaSteps = func(w http.ResponseWriter, r *http.Request) {
	GantaRestrictedGetHelp("children", w, r)
}
