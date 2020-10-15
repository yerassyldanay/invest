package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils"
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

	return utils.ErrorEmailIsNotVerified, err
}

// get list of emails of users, who has connection to the project, by project id
func (e *Email) OnlyGetEmailsHasConnectionToProject(project_id uint64, tx *gorm.DB) ([]Email, error) {
	// get emails of admins & investors
	var emails = []Email{}
	err := tx.Raw(`select distinct e.address from users u join emails e on e.id = u.email_id
		join roles r on r.id = u.role_id left join projects p on u.id = p.offered_by_id 
		where r.name = 'admin' or p.id = ?; `, project_id).Scan(&emails).Error

	if err != nil {
		return emails, err
	}

	// get emails of assigned users
	var spkEmails = []Email{}
	err = tx.Raw(`select distinct e.address from users u join emails e on e.id = u.email_id 
		join projects_users pu on u.id = pu.user_id where project_id = ?; `, project_id).Scan(&spkEmails).Error

	for _, email := range spkEmails {
		emails = append(emails, email)
	}

	return emails, err
}

func (e *Email) OnlyGetEmailsOfSpkUsersAndAdmins(project_id uint64, tx *gorm.DB) ([]Email, error) {
	// get emails of assigned users
	var spkEmails = []Email{}
	err := tx.Raw(`select distinct e.address from users u join emails e on e.id = u.email_id 
		join projects_users pu on u.id = pu.user_id where project_id = ?; `, project_id).Scan(&spkEmails).Error

	// get emails of admins & investors
	var emails = []Email{}
	err = tx.Raw(`select e.* from users u join roles r on r.id = u.role_id ` +
		` join emails e on u.email_id = e.id where r.name = 'admin'; `).Scan(&emails).Error

	for _, email := range spkEmails {
		emails = append(emails, email)
	}

	return emails, err
}


