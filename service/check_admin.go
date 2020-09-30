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
