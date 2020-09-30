package control

import (
	"encoding/json"
	"invest/model"
	"invest/utils"
	"net/http"
)

/*
	201 - created
	400 - bad request
	409 - already in use
	417 - db error
	422 - could not sent message & not stored on db
	500 - internal server error
 */
var Sign_up = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Sign_up"
	var user = model.User{}

	if err := user.ValidateSignUpUser(); err != nil {
		utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " v", err.Error()})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Respond(w, r, utils.Msg{
			Message: utils.ErrorInternalServerError,
			Status:  http.StatusInternalServerError,
			Fname:   fname + " 1",
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	msg := user.Sign_Up()
	utils.Respond(w, r, msg)
}
