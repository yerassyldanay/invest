package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils"

	"net/http"
)

var Create_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Create_project"
	var projectWithFinTable = model.ProjectWithFinanceTables{}

	/*
		only an investor can create a project
	 */
	roleName := service.Get_header_parameter(r, utils.KeyRoleName, "").(string)
	if roleName != utils.RoleInvestor {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " role", "role must be investor. role is " + roleName})
		return
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&projectWithFinTable); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error() })
		return
	}
	defer r.Body.Close()

	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId:   service.Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64),
			RoleName: roleName,
			Lang:     service.Get_header_parameter(r, utils.HeaderContentLanguage, "").(string),
		},
	}

	// logic is inside this func
	var msg = is.Service_create_project(&projectWithFinTable)
	msg.Fname = fname + " ser"

	if msg.ErrMsg == "" {
		msg.Message = utils.NoErrorFineEverthingOk
	}

	utils.Respond(w, r, msg)
}

var Update_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Update_project_by_investor"
	var project = model.Project{}

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.Respond(w, r, utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	project.Lang = r.Header.Get(utils.HeaderContentLanguage)

	msg := project.Update()
	utils.Respond(w, r, msg)
}


