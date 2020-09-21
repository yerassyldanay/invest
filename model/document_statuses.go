package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

func (d Document) Only_update_statuses(documents []Document, tx *gorm.DB) (error) {
	var statusMap = map[string][]uint64{}
	for _, doc := range documents {
		_, ok := statusMap[doc.Status]
		if !ok {
			statusMap[doc.Status] = []uint64{}
		}

		statusMap[doc.Status] = append(statusMap[doc.Status], doc.Id)
	}

	for status, ids := range statusMap {
		if status != utils.ProjectStatusDone && status != utils.ProjectStatusRejected {
			continue
		}

		if err := tx.Raw("update documents set status = ? where project_id = ? and id in (?);", utils.ProjectStatusDone, d.ProjectId, ids).Error;
			err != nil {
				return err
		}
	}

	return nil
}
