package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

/*
	function assumes that email, role & phone id are set
 */
func (c *User) CreateOnly(trans *gorm.DB) error {
	if c.RoleId == 0 || c.EmailId == 0 || c.PhoneId == 0 {
		return errors.New("email, role or phone id is 0")
	}
	return trans.Create(c).Error
}

func (c *User) DeleteOnlyUserById(trans *gorm.DB) error  {
	return trans.Delete(User{}, "id = ?", c.Id).Error
}

/*
	there is no need to create transaction to get info on email
		however, you cannot send another query while having a transaction open
 */
func (c *User) GetByIdPreloaded(trans *gorm.DB) error {
	return trans.Preload("Email").Preload("Phone").Preload("Role").
		Where("id = ?", c.Id).First(c).Error
}

func (c *User) GetByIdOnlyUser(trans *gorm.DB) error {
	return trans.First(c, "id = ?", c.Id).Error
}
