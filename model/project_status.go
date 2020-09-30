package model

import (
	"github.com/jinzhu/gorm"
)

/*
	this is just to change status of the project
 */
func (p *Project) Change_by_project_id_status_to(status string, trans *gorm.DB) (error) {
	//if status != utils.ProjectStatusPendingAdmin && status != utils.ProjectStatusDone &&
	//	status != utils.ProjectStatusRejected && status != utils.ProjectStatusPendingAdmin {
	//	status = utils.ProjectStatusPendingAdmin
	//}

	return trans.Model(&Project{}).Where("id = ?", p.Id).Update("status", status).Error
}

