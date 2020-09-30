package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils"

	"os"
	"path/filepath"
)

// errors for doc validation
var errorInvalidGantaId = errors.New("invalid ganta id")
var errorDocumentInvalidUri = errors.New("invalid / empty uri")

func (d *Document) Validate() error {
	switch {
	case d.ProjectId < 1:
		return errorInvalidProjectId
	case d.GantaId < 1:
		return errorInvalidGantaId
	case d.Uri == "":
		return errorDocumentInvalidUri
	}

	return nil
}

/*
	create a document
 */
func (d *Document) OnlyCreate(trans *gorm.DB) error {
	return  trans.Create(d).Error
}

/*
	add docs to the project by project_id
		at this moment document is already stored on db
 */
func (d *Document) Add_after_validation() (utils.Msg) {
	if err := d.Validate(); err != nil {
		return ReturnInvalidParameters(err.Error())
	}

	/*
		create a document row on db
	 */
	if err := d.OnlyCreate(GetDB()); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	return ReturnNoError()
}

func (d *Document) Remove_document_based_on_responsibility(responsible interface{}) (utils.Msg) {
	var main_query = `select d.* from gantas g join documents d on g.id = d.ganta_id where d.id = ? and g.responsible = ?;`

	if err := GetDB().Raw(main_query, d.Id, responsible).Scan(d).Error; err != nil {
		return ReturnInternalDbError(err.Error())
	}

	// delete file from the path
	var docpath, err = filepath.Abs(".." + "/invest" + d.Uri)
	if err == nil {
		err = os.Remove(docpath)
	}

	fmt.Println("removed file: ", err)

	return ReturnNoError()
}
