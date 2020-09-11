package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
	"time"
)

/*
	note: 2^32 = 4 294 967 296
*/
type User struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`

	Username		string				`json:"username" validate:"required"`
	Password		string				`json:"password"`

	Fio				string				`json:"fio" validate:"required"`
	
	RoleId			uint64				`json:"role_id"`
	Role			Role				`json:"role" gorm:"foreignkey:RoleId"`

	EmailId			uint64				`json:"email_id"`
	Email			Email				`json:"email" gorm:"foreignkey:EmailId"`

	PhoneId			uint64				`json:"phone_id"`
	Phone			Phone				`json:"phone" gorm:"foreignkey:PhoneId"`

	Verified		bool				`json:"verified" gorm:"default:false"`
	Lang			string				`json:"-" gorm:"-"`
	
	OrganizationId		uint64				`json:"organization_id" gorm:"default:0"`
	Organization		Organization		`json:"organization" gorm:"foreignkey:OrganizationId"`

	Blocked			bool				`json:"blocked" gorm:"default:false"`
	Created				time.Time			`json:"created" gorm:"default:now()"`
}

/*
	this returns the name of the table in the database
		gorm must automatically set the name by itself (adding 's' at the end)
		but it is worth to make sure that the name set correctly
*/
func (User) TableName() string {
	return "users"
}

/*
	Hooks are functions that are called before or after creation/querying/updating/deletion.
	If you have defined specified methods for a model, it will be called automatically when creating, updating, querying, deleting, and if any callback returns an error,
	GORM will stop future operations and rollback current transaction.
	The type of hook methods should be func(*gorm.DB) error

	https://gorm.io/docs/hooks.html
 */
//func (c *User) AfterFind(tx *gorm.DB) error {
//	/*
//		after each get method this will set password to ""
//	 */
//	c.Password = ""
//	return nil
//}

var errorDafultUsersAreBeingAltered = errors.New("cannot delete default users")

/*
	cannot delete default users
 */
func (c *User) BeforeDelete(tx *gorm.DB) error {
	if c.Id <= utils.ConstantDefaultNumberOfUsers {
		return errorDafultUsersAreBeingAltered
	}

	return nil
}

/*
	cannot update default user
 */
func (c *User) BeforeUpdate(tx *gorm.DB) error {
	if c.Id <= utils.ConstantDefaultNumberOfUsers {
		return errorDafultUsersAreBeingAltered
	}

	return nil
}


