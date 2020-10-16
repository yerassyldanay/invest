package control

import (
	"encoding/json"
	"invest/model"
	"invest/service"
	"invest/utils"
	"net/http"
)

var Categors_create_read_update_delete = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Categors_create_read_update_delete"
	var msg = utils.Msg{}
	var c = model.Categor{}

	if r.Method != http.MethodGet {
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " json", err.Error()})
			return
		}
		defer r.Body.Close()
	} else {
		var offset = service.Get_query_parameter_str(r, "offset", "0")
		msg = c.Get_all_categors(offset)

		utils.Respond(w, r, msg)
		return
	}

	/*
		only admins can get create & delete
	 */
	roleName := service.Get_header_parameter(r, utils.KeyRoleName, "").(string)
	if roleName != utils.RoleAdmin {
		utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " role", "role is " + roleName})
		return
	}

	switch r.Method {

	case http.MethodPost:
		msg = c.Create_category()

	case http.MethodDelete:
		msg = c.Delete_category_from_tabe_and_projects()

	default:
		msg =  utils.Msg{}
	}

	msg.SetFname(fname, "r")
	utils.Respond(w, r, msg)
}
