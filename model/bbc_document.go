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
	
	SetDeadline					int64				`json:"set_deadline" gorm:"-"`
	Deadline					time.Time			`json:"deadline" gorm:"default:null"`
	Notified					time.Time			`json:"notified" gorm:"default:now()"`

	Modified					time.Time			`json:"modified" gorm:"default:now()"`
	Created						time.Time			`json:"created" gorm:"default:now()"`

	Status						string				`json:"status" gorm:"default:'new_one'"`
	Step						int					`json:"step" gorm:"default:1"`
	IsAdditional				bool				`json:"is_additional" gorm:"default:false"`

	ProjectId					uint64 				`json:"project_id" gorm:"foreignkey:projects.id"`
	Responsible					string				`json:"responsible" gorm:"manager"`

	Count						int					`json:"-" gorm:"-"`
}

func (Document) TableName() string {
	return "documents"
}

// errors for doc validation
var errorDocumentInvalidUri = errors.New("invalid / empty uri")
var errorDocumentInvalidName = errors.New("invalid document name")
var errorDocumentInvalidDeadline = errors.New("invalid year. it is too large")

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

	if d.Kaz == "" {
		d.Kaz = temp
	}

	if d.Rus == "" {
		d.Rus = temp
	}

	if d.Eng == "" {
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
	case d.Deadline.After(utils.GetCurrentTime().Add(time.Hour * 24 * 365 * 5)):
		// the time difference 5 years
		return errorDocumentInvalidDeadline
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
	d.Created = utils.GetCurrentTime()
	err = trans.Create(d).Error
	return err
}

// save
func (d *Document) OnlySave(tx *gorm.DB) (err error) {
	err = tx.Save(d).Error
	return err
}

// only update uri
func (d *Document) OnlyUpdateUriAndDeadlineByIdAndEmptyUri(tx *gorm.DB) (err error) {
	err = tx.Model(&Document{}).Where("id = ? and uri = ''", d.Id).Updates(map[string]interface{}{
		"uri": d.Uri,
		"status": utils.ProjectStatusNewOne,
		"modified": utils.GetCurrentTime(),
		"deadline": utils.GetCurrentTime(),
	}).Error
	return err
}

// get document
func (d *Document) OnlyGetDocumentById(tx *gorm.DB) (err error) {
	err = tx.First(d, "id = ?", d.Id).Error
	return err
}

// get documents
func (d *Document) OnlyGetDocumentsByProjectId(project_id uint64, tx *gorm.DB) ([]Document, error) {
	var documents = []Document{}
	err := tx.Order("created desc").Find(&documents, "project_id = ?", project_id).Error
	return documents, err
}

// get documents based on steps
func (d *Document) OnlyGetDocumentsByStepsAndProjectId(project_id uint64, steps []interface{}, tx *gorm.DB) ([]Document, error) {
	var documents = []Document{}
	err := tx.Order("created desc").Find(&documents, "step in (?) and project_id = ?", steps, project_id).Error
	return documents, err
}

// delete
func (d *Document) OnlyDeleteById(tx *gorm.DB) (err error) {
	err = tx.Model(&Document{Id: d.Id}).Delete(d, "id = ?", d.Id).Error
	return err
}

// set uri to ''
func (d *Document) OnlyEmptyUriById(tx *gorm.DB) (err error) {
	err = tx.Model(&Document{Id: d.Id}).Updates(map[string]interface{}{
		"uri": "",
		"modified": utils.GetCurrentTime(),
	}).Error
	return err
}

// count the number of documents, which are needed to be uploaded by investor
func (d *Document) OnlyCountNumberOfEmptyDocuments(roleName string, step int, tx *gorm.DB) (int) {
	_ = tx.Raw("select count(*) from documents where uri = '' and project_id = ? and step = ? and responsible = ?;", d.ProjectId, step, roleName).
		Count(&d.Count).Error
	return d.Count
}

// a number of documents with undesirable statuses
// those are: reject & accept
// desirable ones are: accept & new_one
func (d *Document) OnlyCountNumberOfDocumentsWithUndesirableStatus(roleName string, project_id uint64, step int, tx *gorm.DB) (int, error) {
	var count int
	err := tx.Table(d.TableName()).Where("project_id = ? and step = ? and status not in (?) and responsible = ?",
		project_id, step, []string{utils.ProjectStatusNewOne, utils.ProjectStatusAccept}, roleName).Count(&count).Error
	return count, err
}

func (d *Document) OnlyUpdateStatusById(tx *gorm.DB) (err error) {
	err = tx.Model(&Document{Id: d.Id}).Updates(map[string]interface{}{
		"status": d.Status,
		"modified": utils.GetCurrentTime(),
	}).Error
	return err
}

func (d *Document) AreAllValidDocumentIds(ids []uint64, project_id uint64, tx *gorm.DB) (bool) {
	var count int
	err := tx.Table(d.TableName()).
		Where("id in (?) and project_id = ?", ids, project_id).
		Count(&count).Error

	ok := (err == nil) && (count == len(ids))
	return ok
}

// get users responsible for document
func (d *Document) OnlyGetEmptyDocumentsWithComingDeadline() ([]Document, error) {
	documents := []Document{}
	currTime := utils.GetCurrentTime()
	err := GetDB().Raw("select min(id) as id, project_id from documents where deadline between ? and ? and " +
		" (uri = '' or status = 'reject' or status = 'reconsider') group by project_id;",
		currTime.Add(time.Hour * 24 * 3), currTime.Add(time.Hour * 24 * 4)).Scan(&documents).Error
	return documents, err
}
