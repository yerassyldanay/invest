package service

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/message"
)

func (is *InvestService) CheckIsItAdmin() message.Msg {
	if is.RoleName != constants.RoleAdmin {
		return model.ReturnMethodNotAllowed("not admin. your role is " + is.RoleName)
	}

	return message.Msg{}
}

func (is *InvestService) DoesUserHasGivenRole(user_id uint64, roles []string) message.Msg {
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
