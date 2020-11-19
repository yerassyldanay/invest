package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// errors
var errorInvalidUsername = errors.New("username is too long")

/*
	Validate:
		* validate username
		* validate password
		* validate bin
		* validate email address
 */
func (c *User) Validate() error {
	// username
	if len(c.Username) < 8 && len(c.Username) > 30 {
		return errorInvalidUsername
	}

	// password
	if err := OnlyValidatePassword(c.Password); err != nil {
		return err
	}

	// org bin
	if err := OnlyValidateBin(c.Organization.Bin); err != nil {
		return err
	}

	// email address
	if err := OnlyValidateEmailAddress(c.Email.Address); err != nil {
		return err
	}

	return nil
}

/*
	DeleteUserByEmail:
		deletes users, who has signed up, but not verified their emails
 */
func (c *User) DeleteUserByEmail(email Email, tx *gorm.DB) error {
	// check for mistake
	if email.Verified {
		return errors.New("this email address has been verified. cannot delete it")
	}

	// check
	if c.Id == 0 {
		if err := c.OnlyGetByEmailAddress(tx); err != nil {
			return err
		}
	}

	// get user preloaded
	var user = User{Id: c.Id}
	if err := user.OnlyGetByIdPreloaded(tx); err != nil {
		return err
	}

	// delete email
	email.Id = c.EmailId
	if err := email.OnlyDeleteById(tx); err != nil {
		return err
	}

	// delete phone number
	var phone = Phone{Id: c.PhoneId}
	if err := phone.OnlyDeleteById(tx); err != nil {
		return err
	}

	// delete user
	if err := user.OnlyDeleteUserById(tx); err != nil {
		return err
	}

	// done!
	return nil
}