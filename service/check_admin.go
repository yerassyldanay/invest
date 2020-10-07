package service

import (
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Check_is_it_admin() utils.Msg {
	if is.RoleName != utils.RoleAdmin {
		return model.ReturnMethodNotAllowed("not admin. your role is " + is.RoleName)
	}

	return utils.Msg{}
}

func (is *InvestService) Does_user_has_given_role(user_id uint64, roles []string) (utils.Msg) {
	// check whether a user exists (preloaded means gets also role, email & phone)
	var user = model.User{Id: user_id}
	err := user.OnlyGetByIdPreloaded(model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// check that user is either manager or expert
	for _, role := range roles {
		if user.Role.Name == role {
			return model.ReturnNoError()
		}
	}

	return model.ReturnMethodNotAllowed("this user cannot access")
}