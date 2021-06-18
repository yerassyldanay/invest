package service

import (
	"encoding/json"
	"fmt"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/message"
	"time"
)

func (is *InvestService) EmailConfirm(userEmail model.Email) (message.Msg) {

	// get profile by code
	userString, err := model.GetRedis().Get(userEmail.SentCode).Result()
	if err != nil {
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("failed to retrieve user profile data from redis. err: %s", err))
	}

	// unmarshal back
	var newUser = model.User{}
	if err := json.Unmarshal([]byte(userString), &newUser); err != nil {
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("failed to unmarshal profile data. err: %s", err.Error()))
	}
	//HelperPrint(newUser)

	// start transaction
	var trans = model.GetDB().Begin()
	defer func() {
		if trans != nil {trans.Rollback()}
	}()

	// get investor role
	var role = model.Role{}
	if err := trans.First(&role, "name = ?", constants.RoleInvestor).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("failed to get investor role. err: %s", err.Error()))
	}

	// create email
	newUser.Email.Verified = true
	if err := trans.Create(&newUser.Email).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("failed to create email. err: %s", err.Error()))
	}

	// create phone
	newUser.Phone.Verified = true
	if err := trans.Create(&newUser.Phone).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("failed to create phone. err: %s", err.Error()))
	}

	// create organization
	if msg := newUser.Organization.Create_or_get_organization_from_db_by_bin(trans); msg.ErrMsg != "" {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("failed to set organization. err: %s", msg.ErrMsg))
	}

	// create user
	newUser = model.User{
		Id:             0,
		Password:       newUser.Password,
		Fio:            newUser.Fio,
		RoleId:         role.Id,
		EmailId:        newUser.Email.Id,
		PhoneId:        newUser.Phone.Id,
		Verified:       true,
		Lang:           newUser.Lang,
		OrganizationId: newUser.Organization.Id,
		Blocked:        false,
		Created:        time.Now(),
	}
	if err := trans.Create(&newUser).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("failed to commit changes to database. err: %s", err.Error()))
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}
	trans = nil

	//
	model.GetDB().First(&newUser, "id = ?", newUser.Id)
	model.GetDB().First(&newUser.Email, "id = ?", newUser.EmailId)
	model.GetDB().First(&newUser.Phone, "id = ?", newUser.PhoneId)
	model.GetDB().First(&newUser.Organization, "id = ?", newUser.OrganizationId)

	newUser.Password = ""
	return model.ReturnNoErrorWithResponseMessage(map[string]interface{}{
		"user": newUser,
	})
}
