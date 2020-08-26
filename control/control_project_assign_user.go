package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Admin_assign_user_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Admin_assign_user_to_project"
	var pu = model.ProjectsUsers{}

	if err := json.NewDecoder(r.Body).Decode(&pu); err  != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: 	utils.ErrorInvalidParameters,
			Status:  	400,
			Fname:   	fname + " 1",
			ErrMsg:  	err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var errmsg string
	//resp, err := pu.Assign_user_to_project()
	//if err != nil {
	//	errmsg = err.Error()
	//}

	resp, err := pu.Notify_both("")
	if err != nil {
		errmsg = errmsg + " | " + err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: 	resp,
		Status:  	utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   	fname,
		ErrMsg:  	errmsg,
	})
}

/*
	remove user & project relation
 */
var Remove_user_from_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Remove_user_from_project"
	var pu = model.ProjectsUsers{}

	if err := json.NewDecoder(r.Body).Decode(&pu); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var errmsg string
	resp, err := pu.Remove_user_from_project()
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:  fname + " 2",
		ErrMsg:  errmsg,
	})
}


