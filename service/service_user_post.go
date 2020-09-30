package service

import (
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Create_user_based_on_role(new_user *model.User) (msg utils.Msg) {

	if err := new_user.ValidateSpkUser(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	msg = new_user.Create_user_without_check()

	return msg
}
