package service

import (
	"github.com/jinzhu/gorm"
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Update_user_profile(user *model.User) (utils.Msg) {
	/*
		get the user account, which is being modified
	 */
	var trans = model.GetDB().Begin()
	defer func() { if trans != nil {trans.Rollback()} }()

	var tempuser = model.User{
		Id: user.Id,
	}

	// get user info
	if err := tempuser.OnlyGetByIdPreloaded(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// if this is a new phone number
	if tempuser.Phone.Ccode + tempuser.Phone.Number != user.Phone.Ccode + user.Phone.Number {
		// delete phone number
		if err := tempuser.Phone.OnlyDeleteById(trans); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		tempuser.Phone.Ccode = user.Phone.Ccode
		tempuser.Phone.Number = user.Phone.Number

		// validate
		if err := user.Phone.Validate(); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// create a new phone number
		if err := tempuser.Phone.OnlyCreate(trans); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		tempuser.PhoneId = tempuser.Phone.Id
	}

	tempuser.Fio = user.Fio

	// save changes
	if err := tempuser.OnlySave(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	trans = nil
	return model.ReturnNoError()
}

func (is *InvestService) Update_user_password(new_password string) (utils.Msg) {
	var user = model.User{Id: is.UserId}

	// check validity of the password
	// for more info refer to the description of the function below
	if err := model.OnlyValidatePassword(new_password); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// convert to hash
	hashed_password, err := utils.Convert_string_to_hash(new_password)
	if err != nil {
		return model.ReuturnInternalServerError(err.Error())
	}

	// only update password
	err = user.OnlyUpdatePasswordById(hashed_password, model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}

/*
	POST
 */
func (is *InvestService) Create_user_based_on_role(new_user *model.User) (utils.Msg) {

	// validate
	if err := new_user.ValidateSpkUser(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// create
	msg := new_user.Create_user_without_check()

	// send notification
	np := model.NotifyCreateProfile{
		UserId:      new_user.Id,
		User:        *new_user,
		CreatedById: 	is.UserId,
	}

	// handles everything
	select {
	case model.GetMailerQueue().NotificationChannel <- &np:
	default:
	}

	return msg
}

/*
	GET
 */
func (is *InvestService) Get_users_by_roles(roles []string) (utils.Msg) {
	var user = model.User{}

	// get users
	users, err := user.OnlyGetUsersByRolePreloaded(roles, is.Offset, model.GetDB())

	// handle err
	switch {
	case err == gorm.ErrRecordNotFound:
		users = []model.User{}
	case err != nil:
		return model.ReturnInternalDbError(err.Error())
	}

	// convert
	var usersMap = []map[string]interface{}{}
	for i, _ := range users {
		users[i].Password = ""
		usersMap = append(usersMap, model.Struct_to_map(users[i]))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = usersMap

	return model.ReturnNoErrorWithResponseMessage(resp)
}
