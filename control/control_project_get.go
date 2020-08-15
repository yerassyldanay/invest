package control

import (
	"invest/model"
	"invest/utils"
	"net/http"
	"strconv"
)

/*
	admin can get all projects with users assigned to them
 */
var User_project_get_all = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_project_get_all"
	offset := Get_query_parameter_str(r, "offset", "0")
	id, _ := strconv.ParseInt(r.Header.Get(utils.KeyId), 0, 16)

	var user = model.User{
		Id: uint64(id),
	}

	var errmsg string
	resp, err := user.Get_all_projects(offset)
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname + " 1",
		ErrMsg:  errmsg,
	})
}

/*
	user (manager, lawyer, financier or others) can get ptojects that have been assigned to them
 */
var User_project_get_own = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_project_get_all"

	offset := Get_query_parameter_str(r, "offset", "0")
	id, _ := strconv.ParseInt(r.Header.Get(utils.KeyId), 0, 16)

	var user = model.User{
		Id: uint64(id),
	}

	var errmsg string
	resp, err := user.Get_own_projects(offset)
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname + " 1",
		ErrMsg:  errmsg,
	})
}
