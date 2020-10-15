package service

import (
	"github.com/jinzhu/gorm"
	"invest/model"
	"invest/utils"
	"time"
)

/*
	signup for investors
		201 - created
		400 - bad request
		409 - already in use
		417 - db error
		422 - could not sent message & not stored on db
*/
func (is *InvestService) SignUp(c model.User) (utils.Msg) {

	// validate
	if err := c.Validate(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// starting a transaction, which is be rolled back in case of an error
	trans := model.GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() } }()

	// there are four cases
	var email = model.Email{Address: c.Email.Address}
	err := email.OnlyGetByAddress(model.GetDB())

	switch {
	case err == gorm.ErrRecordNotFound:
		// pass everything is fine
		// there is no such email address
	case err != nil:
		// unexpected error
		return model.ReturnInternalDbError("err: " + err.Error())
	case email.Verified: // err != nil
		// already in use
		return model.ReturnInvalidParameters("this email is already in use")
	default: // found & not verified
		// delete email address
		if err := c.DeleteUserByEmail(email, trans); err != nil {
			return model.ReturnInternalDbError("delete: " + err.Error())
		}
	}

	// convert password to hash
	hashed, err := utils.Convert_string_to_hash(c.Password)
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// get role id by name (of the role)
	var role = model.Role{Name: utils.RoleInvestor}
	if err := role.OnlyGetByName(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// set fields
	c.RoleId = c.Role.Id
	c.Password = hashed

	// create organization or get old one
	c.Organization.Lang = c.Lang
	c.Organization.Create_or_get_organization_from_db_by_bin(model.GetDB())
	c.OrganizationId = c.Organization.Id

	// these code and link will be sent to the user
	scode := utils.Generate_Random_Number(utils.MaxNumberOfDigitsSentByEmail)

	// store the email, phone & get ids
	c.Email.SentCode = scode
	c.Email.Deadline = utils.GetCurrentTime().Add(time.Hour * 24)

	// create email
	if err := c.Email.OnlyCreate(trans); err != nil {
		//trans.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}
	c.EmailId = c.Email.Id

	// create phone
	if err := c.Phone.OnlyCreate(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}
	c.PhoneId = c.Phone.Id

	// create user
	if err := c.OnlyCreate(trans); err != nil {
		return model.ReturnFailedToCreateAnAccount(err.Error())
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// send notification
	nc := model.NotifyCode{
		Code:    scode,
	}

	select {
	case model.GetMailerQueue().NotificationChannel <- &nc:
	default:
	}

	return model.ReturnNoError()
}
