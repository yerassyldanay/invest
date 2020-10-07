package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"invest/utils"
	"time"
)

type Document struct {
	Id							uint64				`json:"id" gorm:"AUTO_INCREMENT; primary_key"`

	Kaz							string				`json:"kaz" gorm:"not null"`
	Rus							string				`json:"rus" gorm:"not null"`
	Eng							string				`json:"eng" gorm:"not null"`
	
	Uri							string				`json:"uri" gorm:"default:''"`
	
	Modified					time.Time			`json:"modified" gorm:"default:now()"`
	
	SetDeadline					int64					`json:"set_deadline" gorm:"-"`
	Deadline					time.Time			`json:"deadline" gorm:"default:null"`
	Notified					time.Time			`json:"notified" gorm:"default:now()"`

	Step						int					`json:"step" gorm:"default:1"`
	IsAdditional				bool				`json:"is_additional" gorm:"default:false"`

	ProjectId					uint64 				`json:"project_id"`
	Responsible					string				`json:"responsible" gorm:"manager"`
}

func (Document) TableName() string {
	return "documents"
}

// errors for doc validation
var errorDocumentInvalidUri = errors.New("invalid / empty uri")
var errorDocumentInvalidName = errors.New("invalid document name")

// prettify name of the document
func (d *Document) PrettifyName() (err error) {
	var temp = ""
	for _, name := range []string{d.Kaz, d.Rus, d.Eng} {
		if name != "" {
			temp = name
			break
		}
	}

	if temp == "" {
		return errorDocumentInvalidName
	}

	switch {
	case d.Kaz == "":
		d.Kaz = temp
	case d.Rus == "":
		d.Rus = temp
	case d.Eng == "":
		d.Eng = temp
	}

	return nil
}

// validate the info on a document
func (d *Document) Validate() error {
	switch {
	case d.PrettifyName() != nil:
		return errorDocumentInvalidName
	case d.ProjectId < 1:
		return errorInvalidProjectId
	//case d.Uri == "":
	//	return errorDocumentInvalidUri
	}

	switch {
	case d.Step > 2:
		d.Step = 2
	case d.Step < 1:
		d.Step = 1
	}

	if d.Responsible != utils.RoleSpk && d.Responsible != utils.RoleInvestor {
		d.Responsible = utils.RoleSpk
	}

	return nil
}

// create a document
func (d *Document) OnlyCreate(trans *gorm.DB) (err error) {
	err = trans.Create(d).Error
	return err
}

// save
func (d *Document) OnlySave(tx *gorm.DB) (err error) {
	err = tx.Save(d).Error
	return err
}

// only update uri
func (d *Document) OnlyUpdateUriById(tx *gorm.DB) (err error) {
	err = tx.Model(&Document{Id: d.Id}).Update("uri", d.Uri).Error
	return err
}

// get document
func (d *Document) OnlyGetDocumentById(tx *gorm.DB) (err error) {
	err = tx.First(d, "id = ?", d.Id).Error
	return err
}

// get documents
func (d *Document) OnlyGetDocumentsByProjectId(project_id uint64, tx *gorm.DB) (documents []Document, err error) {
	err = tx.Find(&documents, "project_id = ?", project_id).Error
	return documents, err
}

// get documents based on steps
func (d *Document) OnlyGetDocumentsByStepsAndProjectId(project_id uint64, steps []interface{}, tx *gorm.DB) (documents []Document, err error) {
	err = tx.Find(&documents, "step in (?) and project_id = ?", steps, project_id).Error
	return documents, err
}

// delete
func (d *Document) OnlyDeleteById(tx *gorm.DB) (err error) {
	err = tx.Model(&Document{Id: d.Id}).Delete(d, "id = ?", d.Id).Error
	return err
}

// set uri to ''
func (d *Document) OnlyEmptyUriById(tx *gorm.DB) (err error) {
	err = tx.Model(&Document{Id: d.Id}).Update("uri", "").Error
	return err
}

// count the number of documents, which are needed to be uploaded by investor
func (d *Document) OnlyCountNumberOfEmptyDocuments(roleName string, step int, tx *gorm.DB) (count int) {
	_ = tx.Raw("select * from documents where project_id = ? and step = ? and responsible = ?;", d.ProjectId, step, roleName).
		Count(&count).Error
	return count
}

func (d *Document) OnlyCountNumberOfNotAcceptedDocuments(project_id uint64, step int, tx *gorm.DB) (count int) {
	main_query := " select * from documents d join document_user_statuses dus " +
		" on d.id = dus.document_id where project_id = ? and step = ? and status != ?; "
	_ = tx.Raw(main_query, project_id, step, utils.ProjectStatusAccept).Count(&count).Error
	return count
}
