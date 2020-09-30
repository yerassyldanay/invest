package control

import (
	"encoding/json"
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
			utils.Msg{
				Message: utils.ErrorInvalidParameters, Status:  400, Fname: fname + " 1", ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	var msg utils.Msg
	msg = sis.Sign_in()
	msg.Fname = fname + " 3"

	/*
		request header will carry auth token
	*/
	r.Header.Set(utils.HeaderAuthorization, sis.TokenCompound)

	utils.Respond(w, r, msg)
}
