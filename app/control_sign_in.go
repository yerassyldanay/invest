package app

import (
	"encoding/json"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"

	"net/http"
)

/*
	key:
	password
	role
*/
var Sign_in = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Sign_in"
	var sis = model.SignIn{}

	//fmt.Println("r.Body: ", r.Header, r.URL)
	if err := json.NewDecoder(r.Body).Decode(&sis); err != nil {
		message.Respond(w, r,
			message.Msg{
				Message: errormsg.ErrorInvalidParameters, Status:  400, Fname: fname + " 1", ErrMsg:  err.Error(),
		})
		return
	}
	defer r.Body.Close()

	// check email address

	var msg message.Msg
	msg = sis.Sign_in()
	msg.Fname = fname + " 3"

	/*
		request header will carry middleware token
	*/
	r.Header.Set(constants.HeaderAuthorization, sis.TokenCompound)
	//fmt.Println(sis.TokenCompound)
	message.Respond(w, r, msg)
}
