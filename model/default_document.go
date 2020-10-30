package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
	"time"
)

func (d *Document) Set_default_parameters(project_id uint64, step int, investorHasDays time.Duration) {
	d.ProjectId = project_id
	d.Modified = utils.GetCurrentTime()

	// either 1 or 2
	d.Step = step
}

// create documents
func (d *Document) Create_default_documents(project_id uint64, tx *gorm.DB) (utils.Msg) {
	// an investor has 3 days to upload first list of documents
	var firstDeadline = time.Duration(3)

	// investor
	var secondDeadline = time.Duration(376 + 3)

	// prepare the first list
	var err error
	for _, doc := range DefaultDocuments1 {
		doc.Set_default_parameters(project_id, 1, firstDeadline)
		if err = doc.OnlyCreate(tx); err != nil {
			return ReturnInternalDbError(err.Error())
		}
	}

	// prepare the second list
	for _, doc := range DefaultDocuments2 {
		doc.Set_default_parameters(project_id, 2, secondDeadline)
		if err = doc.OnlyCreate(tx); err != nil {
			return ReturnInternalDbError(err.Error())
		}
	}

	return ReturnNoError()
}
