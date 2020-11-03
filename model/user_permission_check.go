package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils/errormsg"
	"invest/utils/message"
	"net/http"
)

/*
	this method checks whether a user has such permission
		provided:
			1. user id
			2. permission name

	statuses: 200, 405
 */
func (up *UserPermission) Check_db_whether_this_user_has_such_a_permission() (message.Msg) {
	var fname = "Check_on_db_whether_this_user_has_such_a_permission"
	if up.Permission == "all" {
		return message.Msg{
			Status: http.StatusOK,
		}
	}

	//fmt.Println("up: ", up.UserId, up.Permission)
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
	if err := GetDB().Raw(main_query, up.UserId, up.Permission).Scan(&user).Error; err == gorm.ErrRecordNotFound || user.Id == 0 {
		return message.Msg{
			Message: errormsg.ErrorMethodNotAllowed, Status:  http.StatusFailedDependency,  Fname:   fname,
			ErrMsg:  "user has not got such permission or invalid parameters have been provided",
		}
	}

	return message.Msg{
		Status: http.StatusOK,
	}
}

/*
	statuses: 200, 406, 417
 */
func (up *UserPermission) Check_db_whether_this_user_account_is_confirmed() (message.Msg) {
	var user = User{}
	if err := GetDB().Table(User{}.TableName()).Where("id = ?", up.UserId).First(&user).Error; err != nil {
		return message.Msg{
			Status:  http.StatusExpectationFailed,
		}
	}

	if !user.Verified {
		return message.Msg{
			errormsg.ErrorEmailIsNotVerified, http.StatusLocked, "", "the account is not confirmed",
		}
	}

	return message.Msg{Status: http.StatusOK}
}
