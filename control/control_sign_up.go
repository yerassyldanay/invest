package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

var Investor_sign_up = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Investor_sign_up"
	var user = model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: utils.ErrorInternalServerError,
			Status:  http.StatusBadRequest,
			Fname:   fname,
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var errmsg string
	resp, err := user.Sign_Up()
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", http.StatusOK, http.StatusBadRequest).(int),
		Fname:   fname,
		ErrMsg:  errmsg,
	})
}
