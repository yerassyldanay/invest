package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
)

type ElementGetFullInfoOfThisUser struct {
	Key   string
	Value string
}

func (c *User) GetFullInfoOfThisUser(e ElementGetFullInfoOfThisUser) message.Msg {
	//HelperPrint(e)

	var err error
	switch e.Key {
	case "email":
		var email = Email{}
		err = GetDB().First(&email, "address = ?", e.Value).Error
		if err != nil {
			fmt.Println(fmt.Errorf("[ERROR] failed to fetch email. err: %v", err))
		}
		err = GetDB().First(c, "email_id = ?", email.Id).Error
	case "phone":
		var phone = Phone{}
		err = GetDB().First(&phone, "ccode || number = ?", e.Value).Error
		if err != nil {
			fmt.Println(fmt.Errorf("[ERROR] failed to fetch phone. err: %v", err))
		}
		err = GetDB().First(c, "phone_id = ?", phone.Id).Error
	default:
		err = GetDB().First(c, "id = ?", c.Id).Error
	}

	if err == gorm.ErrRecordNotFound {
		return message.Msg{errormsg.ErrorNoSuchUser, 404, "", err.Error()}
	}

	_ = GetDB().First(&c.Phone, "id = ?", c.PhoneId)
	_ = GetDB().First(&c.Email, "id = ?", c.EmailId)
	_ = GetDB().First(&c.Role, "id = ?", c.RoleId)
	_ = GetDB().First(&c.Organization, "id = ?", c.OrganizationId)

	return ReturnNoErrorWithResponseMessage(map[string]interface{}{
		"user": *c,
	})
}

func (c *User) GetFullInfoOfThisUserWithoutPasswordById() message.Msg {
	_ = GetDB().First(c, "id = ?", c.Id).Error
	_ = GetDB().First(&c.Phone, "id = ?", c.PhoneId)
	_ = GetDB().First(&c.Email, "id = ?", c.EmailId)
	_ = GetDB().First(&c.Role, "id = ?", c.RoleId)
	_ = GetDB().First(&c.Organization, "id = ?", c.OrganizationId)

	c.Password = ""
	return ReturnNoErrorWithResponseMessage(map[string]interface{}{
		"user": *c,
	})
}

/*
	Get Admins
*/
func (c *User) Get_admins() message.Msg {
	users, err := c.OnlyGetPreloadedUsersByRole(constants.RoleAdmin, GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var usersMap = make([]map[string]interface{}, len(users))
	for _, user := range users {
		usersMap = append(usersMap, Struct_to_map(user))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = usersMap

	return ReturnNoErrorWithResponseMessage(resp)
}

// get users by role (email, phone & role preloaded)
func (c *User) OnlyGetPreloadedUsersByRole(role string, tx *gorm.DB) (users []User, err error) {
	err = tx.Preload("Phone").Preload("Email").Preload("Role").
		Find(&users, "role_id = (select id from roles where name = ? limit 1)", role).Error
	return users, err
}
