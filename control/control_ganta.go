package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Ganta_add_new_step = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_add_new_step"
	var ganta = model.Ganta{}
	if err := json.NewDecoder(r.Body).Decode(&ganta); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname,
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var errmsg string
	resp, err := ganta.Add_new_step()
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname,
		ErrMsg:  errmsg,
	})
}

var Ganta_get_all_steps_by_project_id = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Ganta_get_all_steps_by_project_id"
	var ganta = model.Ganta{
		ProjectId: uint64(Get_query_parameter_int(r, "project_id", 0)),
	}

	var errmsg string
	resp, err := ganta.Get_ganta_by_project_id()
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname,
		ErrMsg:  errmsg,
	})
}

var Update_or_remove_ganta_step = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Update_ganta_step"
	var gu = model.GantaUpDate{}

	if err := json.NewDecoder(r.Body).Decode(&gu); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname,
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var errmsg string
	resp, err := gu.Update_step_start_thus_others()
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname,
		ErrMsg:  errmsg,
	})
}

/*
	delete ganta step
 */
var Delete_ganta_step = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Delete_ganta_step"
	var gu = model.Ganta{}

	if err := json.NewDecoder(r.Body).Decode(&gu); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  400,
			Fname:   fname,
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var errmsg string
	resp, err := gu.Delete_ganta_step()
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname,
		ErrMsg:  errmsg,
	})
}
