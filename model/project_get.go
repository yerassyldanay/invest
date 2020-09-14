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
