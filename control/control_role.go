package control

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"invest/model"
	"invest/service"
	"invest/utils/constants"
	"invest/utils/errormsg"
	"invest/utils/message"

	"net/http"
	"strconv"
)

/*
	delete - ../role/
 */
var Role_create_update_add_and_remove_permissions = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Role_create_update_add_and_delete_permissions_remove"
	var msg message.Msg
	var role = model.Role{}

	/*
		get
	 */
	switch r.Method {
	case http.MethodGet:
		offset := service.Get_query_parameter_str(r, constants.KeyOffset, "0")
		msg = role.Get_roles(offset)
	default:
		if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
			message.Respond(w, r, message.Msg{
				Message: errormsg.ErrorInvalidParameters,
				Status:  http.StatusBadRequest,
				Fname:   fname + " 1",
				ErrMsg:  err.Error(),
			})
			return
		}
		defer r.Body.Close()

		/*
			put and post
		 */
		if r.Method == http.MethodPut {
			msg = role.Update_role_name_description_and_permissions()
		} else if r.Method == http.MethodPost {
			msg = role.Create_a_role_with_permissions()
		} else {
			/*
				if none of these methods, then
			 */
			msg = message.Msg{
				Message: errormsg.ErrorMethodNotAllowed,
				Status:  http.StatusMethodNotAllowed,
				Fname:   fname + " 2",
				ErrMsg:  "this method is not supported",
			}
		}
	}

	message.Respond(w, r, msg)
}

/*
	../role/{role_id}
 */
var Role_delete_or_get_with_role_id = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Role_delete_or_get_with_role_id"
	var msg message.Msg
	var role = model.Role{}

	var vars = mux.Vars(r)
	role.Id, _ = strconv.ParseUint(vars["role_id"], 10, 64)

	switch r.Method {
	case http.MethodDelete:
		msg = role.Delete_this_role()
	case http.MethodGet:
		msg = role.Get_role_info()
	default:
		msg = message.Msg{
			Message: errormsg.ErrorMethodNotAllowed,
			Status:  http.StatusMethodNotAllowed,
			Fname:   fname + " 2",
			ErrMsg:  "this method is not supported",
		}
	}

	message.Respond(w, r, msg)
}

var Role_add_and_remove_permissions = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Role_add_and_remove_permissions"
	var msg = message.Msg{errormsg.ErrorInvalidParameters, 400, fname + " 1", ""}
	var role = model.Role{}

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		msg.ErrMsg = err.Error()
		message.Respond(w, r, msg)
		return
	}
	defer r.Body.Close()

	switch r.Method {
	case http.MethodDelete:
		/*
			remove permissions
		 */
		msg = role.Remove_a_list_of_permissions()

	case http.MethodPost:
		/*
			add permissions
		 */
		msg = role.Add_a_list_of_permissions()

	default:
		/*
			not supported
		 */
		msg = message.Msg{errormsg.ErrorMethodNotAllowed, 405, "", "method is not allowed. role/permissions"}

	}

	message.Respond(w, r, msg)
}
