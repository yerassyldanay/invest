package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils/constants"
	"invest/utils/errormsg"
	"time"
)

/*

 */
type Email struct {
	Id						uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`

	Address					string				`json:"address" gorm:"UNIQUE; size: 255" validate:"required,email"`
	Verified				bool				`json:"verified"  gorm:"default:false"`

	SentCode				string				`json:"sent_code" gorm:"size: 10"`
	Deadline				time.Time			`json:"deadline" gorm:"default:null"`
}

func (Email) TableName() string {
	return "emails"
}

/*

 */
func (e *Email) OnlyDeleteById(trans *gorm.DB) error {
	return trans.Delete(Email{}, "id = ?", e.Id).Error
}

/*
	this function creates an email address
		other fields must be set
*/
func (e *Email) OnlyCreate(trans *gorm.DB) error {
	return trans.Create(e).Error
}

/*
	create new email with hash & code
*/
func (e *Email) CreateEmailWithHashAfterValidation(trans *gorm.DB) error {
	if e.SentCode == "" {
		return errors.New("code is empty")
	}
	return trans.Create(e).Error
}

/*
	pay attention to transaction:
		refer to documentation
*/
func (e *Email) OnlyGetById(trans *gorm.DB) (err error) {
	err = trans.First(e, "id = ?", e.Id).Error
	return err
}

func (e *Email) OnlyGetByAddress(tx *gorm.DB) (err error) {
	err = tx.First(e, "address = ?", e.Address).Error
	return err
}

// get the one, which is not confirmed yet
func (e *Email) OnlyGetNotConfirmedOne(tx *gorm.DB) (error) {
	err := tx.Raw("select * from emails where sent_code != '' and verified = false limit 1;").Scan(e).Error
	return err
}

// free search
func (e *Email) OnlyGetByCode(value string, tx *gorm.DB) (error) {
	err := tx.First(e, "sent_code = ?", value).Error
	return err
}

// free
func (e *Email) OnlyFreeUpAfterConfirmation() {
	e.SentCode = ""
	e.Deadline = time.Time{}
	e.Verified = true
}

// update
func (e *Email) OnlyUpdateAfterConfirmation(tx *gorm.DB) (bool) {
	count := tx.Exec("update emails set sent_code = '', verified = true where address = ?;", e.Address).RowsAffected
	fmt.Println("confirm email, rows affected ", count)
	return true
}

func (e *Email) OnlySave (tx *gorm.DB) (error) {
	return tx.Save(e).Error
}

/*
	is it verified?
 */
func (e *Email) IsVerified() (map[string]interface{}, error) {
	var err error
	if err = GetDB().Model(&Email{}).Where("address", e.Address).First(e).Error; err == nil {
		if e.Verified {
			return nil, nil
		}
	}

	return errormsg.ErrorEmailIsNotVerified, err
}

// get list of emails of users, who has connection to the project, by project id
func (e *Email) OnlyGetEmailsHasConnectionToProject(project_id uint64, tx *gorm.DB) ([]Email, error) {
	// get emails of managers and experts
	emailsSpk, err := e.OnlyGetEmailsOfSpkUsersExceptAdmins(project_id, tx)
	if err != nil {
		return emailsSpk, err
	}

	// get emails of admins
	emailsAdmins, err := e.OnlyGetAllEmailsByRole(constants.RoleAdmin, tx)
	if err != nil {
		return emailsAdmins, err
	}

	// get email of an investor
	if err := e.OnlyGetEmailOfInvestorByProjectId(project_id, tx); err != nil {
		return []Email{}, err
	}

	// gather in one place
	emailsSpk = append(emailsSpk, emailsAdmins...)
	emailsSpk = append(emailsSpk, *e)

	return emailsSpk, nil
}

// get managers & all experts
func (e *Email) OnlyGetEmailsOfSpkUsersExceptAdmins(project_id uint64, tx *gorm.DB) ([]Email, error) {
	emails := []Email{}
	main_query := `select distinct e.address from projects_users pu
		join users u on pu.user_id = u.id
		join emails e on u.email_id = e.id where project_id = ?;`

	err := tx.Raw(main_query, project_id).Scan(&emails).Error
	return emails, err
}

// get email(s) of manager(s)
func (e *Email) OnlyGetEmailOfManagerByProjectId(project_id uint64, tx *gorm.DB) ([]Email, error) {
	emails := []Email{}
	main_query := `select distinct e.address from projects_users pu
		join users u on pu.user_id = u.id join roles r on u.role_id = r.id
		join emails e on u.email_id = e.id where pu.project_id = ? and r.name = ?;`

	err := tx.Raw(main_query, project_id, constants.RoleManager).Scan(&emails).Error
	return emails, err
}

// get email of investor
func (e *Email) OnlyGetEmailOfInvestorByProjectId(project_id uint64, tx *gorm.DB) (error) {
	main_query := `select e.* from projects p ` +
		` join users u on p.offered_by_id = u.id` +
		` join emails e on u.email_id = e.id where p.id = ? limit 1; `
	err := tx.Raw(main_query, project_id).Scan(e).Error

	return err
}

// get spk users (manager & all experts) and all admins
func (e *Email) OnlyGetEmailsOfSpkUsersAndAdmins(project_id uint64, tx *gorm.DB) ([]Email, error) {

	// get emails of all admins
	emailsOfAdmins, _ := e.OnlyGetAllEmailsByRole(constants.RoleAdmin, GetDB())

	// get emails of all experts
	emailsOfSpkUsers, _ := e.OnlyGetEmailsOfSpkUsersExceptAdmins(project_id, GetDB())
	emailsOfAdmins = append(emailsOfAdmins, emailsOfSpkUsers...)

	return emailsOfAdmins, nil
}

// get all admins
func (e *Email) OnlyGetAllEmailsByRole(role string, tx *gorm.DB) ([]Email, error) {
	var emails = []Email{}
	err := tx.Raw(`select e.* from users u join roles r on r.id = u.role_id ` +
		` join emails e on u.email_id = e.id where r.name = ?; `, role).Scan(&emails).Error

	return emails, err
}
