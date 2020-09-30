package service

import (
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Get_own_emails_by_project_id_after_check(project_id uint64, offset interface{}) (utils.Msg) {
	var user = model.User{
		Id: is.UserId,
	}

	/*
		get email address of the user
	 */
	_ = user.OnlyGetByIdPreloaded(model.GetDB())

	/*
		prepare sendgrid message
	 */
	var sms = model.SendgridMessageStore{
		To:    user.Email.Address,
		ProjectId: project_id,
	}

	msg := sms.Get_messages_by_project_id(offset)
	return msg
}
