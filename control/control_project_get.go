package control

import (
	"errors"
	"invest/model"
	"invest/service"
	"invest/utils"

	"net/http"
	"strconv"
)

/*
	admin can get all projects with users assigned to them
 */
var Get_all_user_projects = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_project_get_all"
	offset := service.Get_query_parameter_str(r, "offset", "0")

	var project = model.Project{}
	msg := project.Get_all_after_preload(offset)

	msg.Fname = fname + " 1"
	utils.Respond(w, r, msg)
}

/*
	user (manager, lawyer, financier or others) can get ptojects that have been assigned to them
 */
var Get_own_projects = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_project_get_all"

	offset := service.Get_query_parameter_str(r, "offset", "0")
	id, _ := strconv.ParseInt(r.Header.Get(utils.KeyId), 0, 16)

	var user = model.User{
		Id: uint64(id),
	}

	var errmsg string
	resp, err := user.Get_own_projects(offset)
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname + " 1",
		ErrMsg:  errmsg,
	})
}

var User_get_projects_info_grouped_by_statuses = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_get_projects_info_grouped_by_statuses"
	var project = model.Project{}

	msg := project.Get_stats_on_projects_grouped_by_statuses()
	msg.Fname = fname + " 1"

	//fmt.Println("r.Header.Get(utils.KeyId): ", r.Header.Get(utils.KeyId))

	utils.Respond(w, r, msg)
}


var Get_project_by_project_id = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Get_project_by_project_id"
	var id = service.Get_header_parameter(r, utils.KeyId, uint64(0)).(uint64)

	var project_id = service.Get_query_parameter_uint64(r, "project_id", 0)
	var project = model.Project{Id: project_id}
	var err error = nil

	/*
		Permission
	 */
	roleName := service.Get_header_parameter(r, utils.KeyRoleName, "").(string)
	if roleName == utils.RoleManager || roleName == utils.RoleExpert {
		err = project.OnlyCheckUserByProjectAndUserId(project_id, id, model.GetDB())
	} else if roleName == utils.RoleInvestor {
		project.OfferedById = id
		err = project.OnlyCheckInvestorByProjectAndInvestorId(model.GetDB())
	} else if roleName == utils.RoleAdmin {
		// green light is on for admins
	} else {
		err = errors.New("your role is " + roleName)
	}

	// err means there is something wrong
	if err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, "", err.Error()})
		return
	}

	/*
		getting project here
	 */
	msg := project.Get_this_project_by_project_id()
	msg.Fname = fname + " 1"

	utils.Respond(w, r, msg)
}