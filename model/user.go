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
	if err := Validate_password(c.Password); err != nil {
		return err
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

// save changes
func (c *User) OnlySave(trans *gorm.DB) error {
	return trans.Save(c).Error
}

// update password
func (c *User) OnlyUpdatePasswordById(pass string, tx *gorm.DB) (err error) {
	err = tx.Model(&User{Id: c.Id}).Update("password", pass).Error
	return err
}

// get users by roles (manager, expert and so on) - preloaded (add info about email, phone, etc)
func (c *User) OnlyGetUsersByRolePreloaded(roles []string, offset interface{}, tx *gorm.DB) (users []User, err error) {
	err = tx.Preload("Organization").Preload("Role").Preload("Email").Preload("Phone").
		Limit(GetLimit).Offset(offset).
		Find(&users, "role_id in (select id from roles where name in (?))", roles).
		Error
	return users, err
}

func (c *User) DoesOwnThisProjectById(project_id uint64, tx *gorm.DB) (err error) {
	var project = Project{}
	err = tx.First(&project, "id = ? and offered_by_id = ?", project_id, c.Id).Error
	return err
}

func (c *User) IsAssignedToThisProjectById(project_id uint64, tx *gorm.DB) (err error) {
	var pu = ProjectsUsers{}
	err = tx.Find(&pu, "project_id = ? and user_id = ?", project_id, c.Id).Error
	return err
}