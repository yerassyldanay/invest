package control

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"invest/model"
	"invest/utils"
	"net/http"
)

/*
	key - how the username is called
	username
	password
	role
*/
var Sign_in = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Sign_in"
	var sis = model.SignIn{}

	if err := json.NewDecoder(r.Body).Decode(&sis); err != nil {
		utils.Respond(w, r, &utils.Msg{
			Message: map[string]interface{}{
				"eng": "invalid parameters have been provided",
			},
			Status:  http.StatusBadRequest,
			Fname:   fname,
			ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var msg = utils.Msg{
		Message: nil,
		Status: 200,
		Fname:   fname,
		ErrMsg:  "",
	}

	if err := validator.Validate(sis); err != nil {
		msg = utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  http.StatusBadRequest,
			Fname:   fname,
			ErrMsg:  err.Error(),
		}
	} else {
		token, resp, err := sis.Sign_in()

		msg.Message = resp
		if err != nil {
			msg.ErrMsg = err.Error()
		}

		/*
			request header will carry auth token
		*/
		r.Header.Set(utils.HeaderAuth, token)
	}

	utils.Respond(w, r, &msg)
}
