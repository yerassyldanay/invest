package model

import (
	"github.com/jinzhu/gorm"
)

/*
	get project by id
 */
func (p *Project) Get_by_id(trans *gorm.DB) error {
	return trans.Table(Project{}.TableName()).Where("id=?", p.Id).First(p).Error
}

func (p *Project) Get_all_categors_by_project_id(trans *gorm.DB) (err error) {
	err = trans.Raw("select distinct c.* from projects_categors pc join categors c on pc.categor_id = c.id where pc.project_id = ? ;", p.Id).Scan(&p.Categors).Error
	return err
}
