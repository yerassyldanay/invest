package control

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"invest/model"
	"invest/utils"
	"net/http"
)

/*
	only admin can create a user

	paths:
		../administrate/user - POST. admin creates a user
		../user/{which} - PUT - any user can make changes to his own info or change his password
			which = password
			which = info
		../user/ - DELETE - delete user by someone with admin privilages
		../user/?role=manager&offset=0 - GET - get users by role or all users
 */
var User_create_read_update_delete = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Admin_create_user"
	var user = model.User{}

	user.Lang = r.Header.Get(utils.HeaderContentLanguage)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Respond(w, r, &utils.Msg{utils.ErrorInvalidParameters, 400, fname + " 1", err.Error()})
		return
	}
	defer r.Body.Close()

	var errmsg string
	var err error
	var resp = utils.ErrorInternalServerError

	switch r.Method {
	case http.MethodPost:
		/*
			create
		*/
		resp, err = user.Create_user()

	case http.MethodPut:
		/*
			update info (/user/info) | password (/user/password)
				in this function we set user id to id from the session token to update info or password
				thus others with (admin privileges) cannot make changes
		 */
		user.Id = uint64(Get_query_parameter_int(r, utils.KeyId, 0))

		var vars = mux.Vars(r)
		switch vars["which"] {
		case "info":
			resp, err = user.Update_user_info()
		case "password":
			resp, err = user.Update_user_password()
		default:
			resp = utils.ErrorInvalidParameters
			err = errors.New("nor info or password. user_crud")
		}

	case http.MethodGet:
		roles, ok := r.URL.Query()["role"]
		offset := Get_query_parameter_str(r, "offset", "0")
		if !ok && len(roles) > 0 {
			resp, err = user.Get_all_users(offset)
		} else {
			resp, err = user.Get_users_by_roles(roles, offset)
		}

	case http.MethodDelete:

		var vars = mux.Vars(r)
		switch vars["which"] {
		case "delete":
			resp, err = user.Delete_user()
		case "block":
			resp, err = user.Block_unblock_user()
		default:
			resp = utils.ErrorInvalidParameters
			err = errors.New("nor delete or block. user_crud")
		}

	default:
		resp = utils.ErrorMethodNotAllowed
		err = errors.New("user crud. method invalid " + r.Method)
	}
	if err != nil {
		errmsg = err.Error()
	}

	utils.Respond(w, r, &utils.Msg{
		Message: resp,
		Status:  utils.If_condition_then(errmsg == "", 200, 400).(int),
		Fname:   fname + " 2",
		ErrMsg:  errmsg,
	})
}


