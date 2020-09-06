package control

import (
	"invest/model"
	"invest/utils"
	"net/http"
)

/*
	this is to confirm that a user owns this email
 */
var User_email_confirm = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_email_confirm"

	key := model.Get_value_from_query(r, "key")
	hashcode := model.Get_value_from_query(r, "hashcode")

	var errmsg string
	var err  error
	var resp = utils.ErrorInvalidParameters

	if hashcode == "" {
		errmsg = "invalid hashcode"
	} else {
		var email = model.Email{
			SentHash: hashcode,
		}
		resp, err = email.Confirm(key)
		if err != nil {
			errmsg = err.Error()
		}
	}

	utils.Respond(w, r, &utils.Msg{
		Message: 	resp,
		Status:  	utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   	fname,
		ErrMsg:  	errmsg,
	})
}

/*
	after getting code on your phone, you can confirm that you own the phone number
 */
var User_phone_confirm = func(w http.ResponseWriter, r *http.Request) {
	var fname = "User_phone_confirm"
	var p = model.Phone{}

	//if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
	//	utils.Respond(w, r, &utils.Msg{
	//		Message: 	utils.ErrorInternalServerError,
	//		Status:  	500,
	//		Fname:   	fname + " 1",
	//		ErrMsg:  	err.Error(),
	//	})
	//}

	p.Number = model.Get_value_from_query(r, "number")
	p.SentCode = model.Get_value_from_query(r, "sent_code")

	resp, err := p.Confirm()

	var errmsg string
	if err != nil{
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: 	resp,
		Status:  	utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   	fname + " 2",
		ErrMsg:  	errmsg,
	})
}

