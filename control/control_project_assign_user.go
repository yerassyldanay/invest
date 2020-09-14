package control

import (
	"encoding/json"
	"fmt"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Assign_user_to_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Admin_assign_user_to_project"
	var pu = model.ProjectsUsers{}

	lang := r.Header.Get(utils.HeaderContentLanguage)

	if err := json.NewDecoder(r.Body).Decode(&pu); err  != nil {
		utils.Respond(w, r, utils.Msg{
			Message: 	utils.ErrorInvalidParameters,
			Status:  	400,
			Fname:   	fname + " 1",
			ErrMsg:  	err.Error(),
		})
		return
	}
	defer r.Body.Close()

	msg := pu.Assign_user_to_project()
	if msg.ErrMsg == "" {
		fmt.Println(lang)
		//_, _ = pu.Notify_user(lang)
	}

	utils.Respond(w, r, msg)
}

/*
	remove user & project relation
 */
var Remove_user_from_project = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Remove_user_from_project"
	var pu = model.ProjectsUsers{}

	if err := json.NewDecoder(r.Body).Decode(&pu); err != nil {
		utils.Respond(w, r, utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	msg := pu.Remove_user_from_project()
	utils.Respond(w, r, msg)
}

