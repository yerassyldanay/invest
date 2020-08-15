package control

import (
	"invest/model"
	"invest/utils"
	"net/http"
	"strconv"
)

/*
	get info for personal info
 */
var User_get_own_info = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_get_own_info"
	var id, _ = strconv.ParseInt(r.Header.Get(utils.KeyId), 0, 16)

	var user = model.User{
		Id: uint64(id),
	}

	var errmsg string
	resp, err := user.Get_full_info_of_this_user("id")
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
