package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println(tempuser.Phone.Ccode + tempuser.Phone.Number, user.Phone.Ccode + user.Phone.Number)
	if tempuser.Phone.Ccode + tempuser.Phone.Number != user.Phone.Ccode + user.Phone.Number &&
		user.Phone.Number != "" {
		// validate
		if err := user.Phone.Validate(); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// delete phone number
		if err := tempuser.Phone.OnlyDeleteById(trans); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		tempuser.Phone.Ccode = user.Phone.Ccode
		tempuser.Phone.Number = user.Phone.Number

		// create a new phone number
		if err := tempuser.Phone.OnlyCreate(trans); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		tempuser.PhoneId = tempuser.Phone.Id
	}

	a := len(user.Fio)
	_ = a
	if len(user.Fio) >= 8 && len(user.Fio) <= 255 {
		tempuser.Fio = user.Fio
	}

	// save changes
	if err := tempuser.OnlySave(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// update key
	_ = model.Update_sequence_id_thus_avoid_duplicate_primary_key_error(model.GetDB(), "phones")

	trans = nil
	return model.ReturnNoError()
}

func (is *InvestService) Update_user_password(old_password, new_password string) (utils.Msg) {
	var user = model.User{Id: is.UserId}

	// get user info
	if err := user.OnlyGetUserById(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// here we check whether two passwords
	// (a provided password and password on db_create_fake_data) MATCH
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old_password));
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return model.ReturnWrongPassword("password is wrong")
	} else if err != nil {
		return model.ReturnInvalidPassword("password either does not match or invalid")
	}

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

	// notify user
	nnp := model.NotifyNewPassword{
		UserId:         is.UserId,
		RawNewPassword: new_password,
	}

	// this handles all other work
	model.GetMailerQueue().NotificationChannel <- &nnp

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
