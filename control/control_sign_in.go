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

	//fmt.Println("r.Body: ", r.Header, r.URL)
	if err := json.NewDecoder(r.Body).Decode(&sis); err != nil {
		utils.Respond(w, r,
			&utils.Msg{
				Message: utils.ErrorInvalidParameters, Status:  400, Fname: fname + " 1", ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	//fmt.Println("sis: ", sis)

	var msg *utils.Msg
	if err := validator.Validate(sis); err != nil {
		msg = &utils.Msg{
			Message: utils.ErrorInvalidParameters,
			Status:  http.StatusBadRequest,
			Fname:   fname + " 2",
			ErrMsg:  err.Error(),
		}
	} else {
		msg = sis.Sign_in()
		msg.Fname = fname + " 3"

		/*
			request header will carry auth token
		*/
		//fmt.Println("sis.TokenCompound: ", sis.TokenCompound)
		r.Header.Set(utils.HeaderAuthorization, sis.TokenCompound)
	}

	//fmt.Printf("\nmsg: %#v \n\n", msg)

	utils.Respond(w, r, msg)
}
