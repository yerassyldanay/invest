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
	var project = model.Project{}

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.Respond(w, r, utils.Msg{
			Message: 	utils.ErrorInvalidParameters,
			Status:  	400,
			Fname:   	fname + " 1",
			ErrMsg:  	err.Error(),
		})
		return
	}
	defer r.Body.Close()

	project.OfferedById = utils.GetHeader(r, utils.KeyId)
	project.AddInfo.Lang = r.Header.Get(utils.HeaderContentLanguage)

	var msg = service.Service_create_project(&project)
	msg.Fname = fname

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


