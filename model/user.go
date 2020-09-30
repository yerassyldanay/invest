package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// errors
var errorSignUpInvalidUsername = errors.New("invalid username: must contain at least 8 and at most 30 characters (digits & letters)")
var errorSignUpInvalidPassword = errors.New("invalid password: must contain at least 8 and at most 30 characters (digits & letters)")
var errorSignUpInvalidFio = errors.New("invalid fio: must contain at least 5 characters")

func (c *User) ValidateSignUpUser() (error) {
	if err := c.ValidateSpkUser(); err != nil {
		return err
	}

	return nil
}

func (c *User) ValidateSpkUser() (error) {
	// username
	if len(c.Username) < 8 || len(c.Username) > 30 {
		return errorSignUpInvalidUsername
	}

	// password
	if len(c.Password) < 8 || len(c.Password) > 30 {
		return errorSignUpInvalidPassword
	}

	return nil
}

func (c *User) ValidateForUpdate() error {
	if !Validate_password(c.Password, "", "") {
		return errorSignUpInvalidPassword
	}

	if len(c.Fio) < 5 {
		return errorSignUpInvalidFio
	}

	// check email address
	// check phone number

	return nil
}

/*
	function assumes that email, role & phone id are set
 */
func (c *User) OnlyCreate(trans *gorm.DB) error {
	if c.RoleId == 0 || c.EmailId == 0 || c.PhoneId == 0 {
		return errors.New("email, role or phone id is 0")
	}
	return trans.Create(c).Error
}

func (c *User) OnlyDeleteUserById(trans *gorm.DB) error  {
	return trans.Delete(User{}, "id = ?", c.Id).Error
}

/*
	there is no need to create transaction to get info on email
		however, you cannot send another query while having a transaction open
 */
func (c *User) OnlyGetByIdPreloaded(trans *gorm.DB) error {
	return trans.Preload("Email").Preload("Phone").Preload("Role").
		Where("id = ?", c.Id).First(c).Error
}

func (c *User) OnlyGetUserById(trans *gorm.DB) error {
	return trans.First(c, "id = ?", c.Id).Error
}

func (c *User) OnlySave(trans *gorm.DB) error {
	return trans.Save(c).Error
}
