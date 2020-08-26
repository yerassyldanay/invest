package model

import (
	"invest/utils"
	"net/http"
)

/*
	this method checks whether a user has such permission
		provided:
			1. user id
			2. permission name
 */
func (up *UserPermission) Check_on_db_whether_this_user_has_such_a_permission_and_user_account_is_confirmed() (*utils.Msg) {
	var fname = "Check_on_db_whether_this_user_has_such_a_permission"
	if up.Permission == "all" {
		return &utils.Msg{}
	}

	var main_query = `
		select u.* from users u where u.id = $1 and u.role_id in
		   (
			   select rp.role_id
			   from roles_permissions rp
						join permissions p on p.id = rp.permission_id
			   where p.name = $2
		   );
	`

	var user = User{}
	if err := GetDB().Raw(main_query, up.UserId, up.Permission).Scan(&user).Error; err != nil {
		return &utils.Msg{
			Message: utils.ErrorMethodNotAllowed, Status:  http.StatusMethodNotAllowed,  Fname:   fname,
			ErrMsg:  "user has not got such permission or invalid parameters have been provided",
		}
	}

	if !user.Verified {
		return &utils.Msg {
			utils.ErrorEmailIsNotVerified, http.StatusNotAcceptable, "", "the account is not confirmed",
		}
	}

	return &utils.Msg{}
}
