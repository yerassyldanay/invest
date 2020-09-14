package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

/*
	this is just to change status of the project
 */
func (p *Project) Change_by_project_id_status_to(status string, trans *gorm.DB) (error) {
	if status != utils.ProjectStatusInprogress && status != utils.ProjectStatusDone &&
		status != utils.ProjectStatusRejected && status != utils.ProjectStatusNewone {
		status = utils.ProjectStatusInprogress
	}

	return trans.Model(&Project{}).Where("id = ?", p.Id).Update("status", status).Error
}

