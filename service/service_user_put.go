package service

import (
	"github.com/jinzhu/gorm"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/helper"
	"github.com/yerassyldanay/invest/utils/message"
)

func (is *InvestService) UpdateUserProfile(user *model.User) message.Msg {
	// get the user account, which is being modified
	var trans = model.GetDB().Begin()

	// get user
	var oldUserProfile = model.User{}
	if err := trans.First(&oldUserProfile, "id = ?", user.Id).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}

	if (user.Phone.Ccode != oldUserProfile.Phone.Ccode ||
		user.Phone.Number != oldUserProfile.Phone.Number) && (len(user.Phone.Number)*len(user.Phone.Ccode) != 0) {
		// delete old phone
		if err := trans.Delete(&model.Phone{}, "id = ?", oldUserProfile.Phone.Id).Error; err != nil {
			_ = trans.Rollback()
			return model.ReturnInternalDbError(err.Error())
		}

		// create new one
		if err := trans.Create(&user.Phone).Error; err != nil {
			_ = trans.Rollback()
			return model.ReturnInternalDbError(err.Error())
		}

		oldUserProfile.PhoneId = user.Phone.Id
	}

	// update fio
	if len(user.Fio) > 0 {
		oldUserProfile.Fio = user.Fio
	}

	// save data
	oldUserProfile.Phone = user.Phone
	if err := trans.Save(&oldUserProfile).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}

func (is *InvestService) UpdateUserPassword(old_password, new_password string) message.Msg {
	var user = model.User{Id: is.UserId}

	// get user info
	if err := model.GetDB().First(&user, "id = ?", is.UserId).Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	if old_password != user.Password {
		return model.ReturnInvalidPassword("old password is not correct")
	}

	//// here we check whether two passwords
	//// (a provided password and password on db_create_fake_data) MATCH
	//err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old_password));
	//if err == bcrypt.ErrMismatchedHashAndPassword {
	//	return model.ReturnWrongPassword("password is wrong")
	//} else if err != nil {
	//	return model.ReturnInvalidPassword("password either does not match or invalid")
	//}

	//// check validity of the password
	//// for more info refer to the description of the function below
	//if err := model.OnlyValidatePassword(new_password); err != nil {
	//	return model.ReturnInvalidParameters(err.Error())
	//}

	//// convert to hash
	//hashed_password, err := helper.Convert_string_to_hash(new_password)
	//if err != nil {
	//	return model.ReuturnInternalServerError(err.Error())
	//}

	// only update password
	if err := model.GetDB().Model(&model.User{}).Where("id = ?", user.Id).
		Updates(map[string]interface{}{
			"password": new_password,
		}).Error; err != nil {
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

// CreateUserBasedOnRole POST
func (is *InvestService) CreateUserBasedOnRole(newUser *model.User) message.Msg {

	rawPassword := newUser.Password

	// role must be (spk)
	ok := helper.DoesASliceContainElement([]string{constants.RoleExpert, constants.RoleAdmin, constants.RoleManager}, newUser.Role.Name)
	if !ok {
		return model.ReturnMethodNotAllowed("role is invalid")
	}

	// validate password
	if len(newUser.Password) < 8 || len(newUser.Password) > 50 {
		return model.ReturnInvalidPassword("too long password is provided")
	}

	// create user
	if err := newUser.ValidateFioPhoneEmail(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// tx
	tx := model.GetDB().Begin()

	// create phone
	var newPhone = newUser.Phone
	if err := tx.Create(&newPhone).Error; err != nil {
		_ = tx.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}
	newUser.PhoneId = newPhone.Id

	// create email
	var newEmail = newUser.Email
	if err := tx.Create(&newEmail).Error; err != nil {
		_ = tx.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}
	newUser.EmailId = newEmail.Id

	// get role
	var role = model.Role{}
	if err := tx.First(&role, "name = ?", newUser.Role.Name).Error; err != nil {
		_ = tx.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}
	newUser.RoleId = role.Id

	// assign user to all projects if role is expert
	if newUser.Role.Name == constants.RoleExpert {
		pu := model.ProjectsUsers{
			UserId: newUser.Id,
		}
		if err := pu.OnlyAssignExpertToAllProjects(tx); err != nil {
			_ = tx.Rollback()
			return model.ReturnInternalDbError(err.Error())
		}
	}

	// create user
	if err := tx.Create(&newUser).Error; err != nil {
		_ = tx.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := tx.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// get user info
	newUser.Role = role
	newUser.Email = newEmail
	newUser.Phone = newPhone

	// send notification
	np := model.NotifyCreateProfile{
		UserId:      newUser.Id,
		User:        *newUser,
		CreatedById: is.UserId,
		RawPassword: rawPassword,
	}

	// handles everything
	select {
	case model.GetMailerQueue().NotificationChannel <- &np:
	default:
	}

	return model.ReturnSuccessfullyCreated()
}

//
func (is *InvestService) GetUsersByRoles(roles []string) message.Msg {
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

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = usersMap

	return model.ReturnNoErrorWithResponseMessage(resp)
}
