package service

import (
	"invest/model"
	"invest/utils"
	"time"
)

func (is *InvestService) EmailConfirm(userEmail model.Email) (utils.Msg) {

	// start transaction
	var trans = model.GetDB().Begin()
	defer func() {
		if trans != nil {trans.Rollback()}
	}()

	// get email with the same code or hash
	var email = model.Email{Address: userEmail.Address}
	if err := email.OnlyGetByAddress(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// check for validity
	// it might be that email in use
	// or a code sent by a user is invalid
	// or the deadline for an email is before the current date
	if email.Verified || email.Deadline.Before(utils.GetCurrentTime()) ||email.SentCode != userEmail.SentCode {
		return model.ReturnInvalidParameters("code is not valid or email is already in use")
	}

	// get rid of extra values
	email.SentCode = ""
	email.Deadline = time.Time{}
	email.Verified = true

	// update values
	if ok := email.OnlyUpdateAfterConfirmation(trans); !ok {
		return model.ReturnInternalDbError("could not update / confirm email on the level of db")
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}
	trans = nil

	return model.ReturnNoError()
}
