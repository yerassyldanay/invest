package app

import (
	"encoding/json"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

var Categors_create_read_update_delete = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Categors_create_read_update_delete"
	var msg = message.Msg{}
	var c = model.Categor{}

	if r.Method != http.MethodGet {
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			message.Respond(w, r, message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " json", err.Error()})
			return
		}
		defer r.Body.Close()
	} else {
		var offset = service.Get_query_parameter_str(r, "offset", "0")
		msg = c.Get_all_categors(offset)

		message.Respond(w, r, msg)
		return
	}

	/*
		only admins can get create & delete
	 */
	roleName := service.Get_header_parameter(r, constants.KeyRoleName, "").(string)
	if roleName != constants.RoleAdmin {
		message.Respond(w, r, message.Msg{errormsg.ErrorMethodNotAllowed, 405, fname + " role", "role is " + roleName})
		return
	}

	switch r.Method {

	case http.MethodPost:
		msg = c.Create_category()

	case http.MethodDelete:
		msg = c.Delete_category_from_tabe_and_projects()

	default:
		msg =  message.Msg{}
	}

	msg.SetFname(fname, "r")
	message.Respond(w, r, msg)
}
