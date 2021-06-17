package app

import (
	"encoding/json"
	"errors"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"

	"net/http"
)

var Create_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Create_project"
	var projectWithFinTable = model.ProjectWithFinanceTables{}

	/*
		only an investor can create a project
	 */
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	//
	if is.RoleName != constants.RoleInvestor {
		message.Respond(w, r, message.Msg{errormsg.ErrorMethodNotAllowed, 405, fname + " role", "role must be investor. role is " + is.RoleName})
		return
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&projectWithFinTable); err != nil {
		message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " 1", err.Error() })
		return
	}
	defer r.Body.Close()

	// logic is inside this func
	var msg = is.Service_create_project(&projectWithFinTable)

	if msg.Status == 0 {
		msg = model.ReturnNoError()
	}

	msg.Fname = fname + " service"
	message.Respond(w, r, msg)
}

var Update_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Update_project_by_investor"
	var project = model.Project{}

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		message.Respond(w, r, message.Msg{
			Message: errormsg.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	project.Lang = r.Header.Get(constants.HeaderContentLanguage)

	msg := project.Update()
	message.Respond(w, r, msg)
}

var Get_project_by_project_id = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_project_by_project_id"

	// headers
	is := service.InvestService{}
	is.OnlyParseRequest(r)

	// parameters
	var project_id = service.Get_query_parameter_uint64(r, "project_id", 0)
	var project = model.Project{Id: project_id}

	var err error = nil

	/*
		Permission
	*/
	roleName := service.Get_header_parameter(r, constants.KeyRoleName, "").(string)
	if roleName == constants.RoleInvestor {
		project.OfferedById = is.UserId
		err = project.OnlyCheckInvestorByProjectAndInvestorId(model.GetDB())
	} else if roleName == constants.RoleAdmin || roleName == constants.RoleManager || roleName == constants.RoleExpert {
		// green light is on for admins
	} else {
		err = errors.New("your role is " + roleName)
	}

	// err means there is something wrong
	if err != nil {
		OnlyReturnMethodNotAllowed(w, r, err.Error(), fname, "err")
		return
	}

	/*
		getting project here
	*/
	msg := is.Project_get_by_id(project_id)
	msg.SetFname(fname, "msg")

	message.Respond(w, r, msg)
}
