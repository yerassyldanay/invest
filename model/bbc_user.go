package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils"

	"time"
)

/*
	note: 2^32 = 4 294 967 296
*/
type User struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`

	Username		string				`json:"username" gorm:"unique" validate:"required"`
	Password		string				`json:"password"`

	Fio				string				`json:"fio" gorm:"unique" validate:"required"`
	
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

	Statistics			UserStats				`json:"statistics" gorm:"-"`
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

var errorDefultUsersAreBeingAltered = errors.New("cannot delete default users")
var errorDefaultUsersAreBeingUpdated = errors.New("cannot update basic user info")

/*
	cannot delete default users
 */
func (c *User) BeforeDelete(tx *gorm.DB) error {
	if c.Id <= utils.ConstantDefaultNumberOfUsers {
		return errorDefultUsersAreBeingAltered
	}

	return nil
}

/*
	cannot update default user
 */
//func (c *User) BeforeUpdate(tx *gorm.DB) error {
//	if c.Id <= utils.ConstantDefaultNumberOfUsers {
//		return errorDefaultUsersAreBeingUpdated
//	}
//
//	return nil
//}

// errors
var errorSignUpInvalidUsername = errors.New("invalid username: must contain at least 8 and at most 50 characters")
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
	if len(c.Username) < 8 || len(c.Username) > 50 {
		fmt.Println("this is fail", c.Username)
		return errorSignUpInvalidUsername
	}

	// password
	if len(c.Password) < 8 || len(c.Password) > 30 {
		return errorSignUpInvalidPassword
	}

	return nil
}

func (c *User) ValidateForUpdate() error {
	if err := OnlyValidatePassword(c.Password); err != nil {
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
	//if c.RoleId == 0 || c.EmailId == 0 || c.PhoneId == 0 {
	//	return errors.New("email, role or phone id is 0")
	//}
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

// get by email address preloaded
func (c *User) OnlyGetByEmailAddress(tx *gorm.DB) (err error) {
	err = tx.Raw("select u.* from users u join emails e " +
		" on u.email_id = e.id where e.address = ? limit 1;", c.Email.Address).
		Scan(c).Error
	return err
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

// get spk users
func (c *User) OnlyGetSpkUsersByProjectId(project_id uint64, tx *gorm.DB) ([]User, error) {
	var users = []User{}
	err := tx.Raw("select u.* from users u join projects_users pu on u.id = pu.user_id where pu.project_id = ?;", project_id).Scan(&users).Error
	return users, err
}

// get spk users by project id & role name
func (c *User) OnlyGetSpkUsersByProjectIdAndRoleName(project_id uint64, roleName string, tx *gorm.DB) ([]User, error) {
	var users = []User{}
	err := tx.Raw("select distinct u.* from users u join projects_users pu on u.id = pu.user_id " +
		"join roles r on u.role_id = r.id where project_id = ? and r.name = ?;", project_id, roleName).Scan(&users).Error
	return users, err
}

// get spk users by project & role names
func (c *User) OnlyGetSpkUsersByProjectIdAndRoles(project_id uint64, roles []string, tx *gorm.DB) ([]User, error) {
	var users = []User{}
	err := tx.Raw("select distinct u.* from users u join projects_users pu on u.id = pu.user_id " +
		"join roles r on u.role_id = r.id where project_id = ? and r.name in (?);", project_id, roles).Scan(&users).Error
	return users, err
}

// get investor
func (c *User) OnlyGetInvestorByProjectId(project_id uint64, tx *gorm.DB) (error) {
	err := tx.Raw("select u.* from projects p join users u on p.offered_by_id = u.id where p.id = ? limit 1", project_id).Scan(c).Error
	return err
}

// get all users who has connection to project + investor + admins
func (c *User) GetAllUsersAndAdminsByProjectId(project_id uint64, tx *gorm.DB) ([]User, error) {
	// get all admins
	var admins = []User{}
	if admins, err := c.OnlyGetPreloadedUsersByRole(utils.RoleAdmin, tx); err != nil && err != gorm.ErrRecordNotFound {
		return admins, err
	}

	var spkUsers = []User{}
	if spkUsers, err := c.OnlyGetSpkUsersByProjectId(project_id, tx); err != nil && err != gorm.ErrRecordNotFound {
		return spkUsers, err
	}

	var investor = User{}
	if err := investor.OnlyGetInvestorByProjectId(project_id, tx); err != nil {
		return []User{}, err
	}

	admins = append(admins, investor)
	for _, spkUsers := range spkUsers {
		admins = append(admins, spkUsers)
	}

	return admins, nil
}
