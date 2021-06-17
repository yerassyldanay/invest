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
	code, err := model.GetRedis().Get(userEmail.SentCode).Result()
	if err != nil {
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("faield to retrieve user profile data from redis. err: %s", err))
	}

	// unmarshal back
	var newUser = model.User{}
	if err := json.Unmarshal([]byte(code), &newUser); err != nil {
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("faield to unmarshal profile data. err: %s", err.Error()))
	}

	// start transaction
	var trans = model.GetDB().Begin()
	defer func() {
		if trans != nil {trans.Rollback()}
	}()

	// get investor role
	var role = model.Role{}
	if err := trans.First(&role, "name = ?", constants.RoleInvestor).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("faield to get investor role. err: %s", err.Error()))
	}

	// create email
	if err := trans.Create(&newUser.Email).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("faield to create email. err: %s", err.Error()))
	}

	// create phone
	if err := trans.Create(&newUser.Phone).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("faield to create phone. err: %s", err.Error()))
	}

	// create organization
	if msg := newUser.Organization.Create_or_get_organization_from_db_by_bin(trans); msg.ErrMsg != "" {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("faield to set organization. err: %s", msg.ErrMsg))
	}

	// create user
	newUser = model.User{
		Id:             0,
		Password:       newUser.Password,
		Fio:            newUser.Fio,
		RoleId:         role.Id,
		EmailId:        newUser.EmailId,
		PhoneId:        newUser.PhoneId,
		Verified:       true,
		Lang:           newUser.Lang,
		OrganizationId: newUser.Organization.Id,
		Blocked:        false,
		Created:        time.Now(),
	}
	if err := trans.Create(&newUser).Error; err != nil {
		_ = trans.Rollback()
		return model.ReturnFailedToCreateAnAccount(fmt.Sprintf("faield to commit changes to database. err: %s", err.Error()))
	}

	// commit changes
	if err := trans.Commit().Error; err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}
	trans = nil

	return model.ReturnNoError()
}
