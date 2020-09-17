package control

import (
	"encoding/json"
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
		../user/ - DELETE - delete user by someone with admin privileges
		../user/?role=manager&offset=0 - GET - get users by role or all users
 */
var User_create_read_update_delete = func(w http.ResponseWriter, r *http.Request) {
	var fname = "Admin_create_user"
	var user = model.User{}

	user.Lang = r.Header.Get(utils.HeaderContentLanguage)
	//fmt.Println(fname, r.URL, r.Header, r.Method, "\n\n")

	/*
		parse user data
	 */
	if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.Respond(w, r, utils.Msg{utils.ErrorInvalidParameters, 400, fname + " post", err.Error()})
			return
		}
		defer r.Body.Close()
	}

	var msg = utils.Msg{Fname: fname + " 1"}

	switch r.Method {
	case http.MethodPost:
		/*
			create
		*/
		msg = user.Create_user()

	case http.MethodPut:
		/*
			update info (/user/info) | password (/user/password)
				in this function we set user id to id from the session token to update info or password
				thus others with (admin privileges) cannot make changes
		 */
		//user.Id = uint64(Get_query_parameter_int(r, utils.KeyId, 0))

		var vars = mux.Vars(r)
		switch vars["which"] {
		case "info":
			msg = user.Update_user_info_except_for_email_address_and_password_by_user_id()
		case "password":
			msg = user.Update_own_user_password_by_user_id()
		default:
			msg = utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " 4", "crud user an unknown put request"}
		}

	case http.MethodGet:

		roles, ok := r.URL.Query()["role"]
		offset := Get_query_parameter_str(r, "offset", "0")
		if !ok && len(roles) == 0 {
			msg = user.Get_all_users_except_admins(offset)
		} else {
			msg = user.Get_users_by_roles(roles, offset)
		}

	case http.MethodDelete:

		var vars = mux.Vars(r)
		switch vars["which"] {
		case "delete":
			msg = user.Delete_user()
		case "block":
			msg = user.Block_unblock_user()
		default:
			msg = utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " 3", "nor delete or block. user_crud"}
		}

	default:
		msg = utils.Msg{utils.ErrorMethodNotAllowed, 405, fname + " 2", "user crud. method invalid " + r.Method}
	}

	utils.Respond(w, r, msg)
}


