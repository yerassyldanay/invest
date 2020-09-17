package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

func (p *Project) Unmarshal_info() error {
	return json.Unmarshal([]byte(p.Info), &p.InfoSent)
}

/*
	get project by id
 */
func (p *Project) Get_by_id(trans *gorm.DB) error {
	return trans.First(p, "id=?", p.Id).Error
}

func (p *Project) Get_all_categors_by_project_id(trans *gorm.DB) (err error) {
	err = trans.Raw("select distinct c.* from projects_categors pc join categors c on pc.categor_id = c.id where pc.project_id = ? ;", p.Id).Scan(&p.Categors).Error
	return err
}

func (p *Project) Get_only_assigned_users_to_project(trans *gorm.DB) (err error) {
	return GetDB().Preload("Email").Preload("Role").Omit("password, created").Find(&p.Users, "id in (select user_id from projects_users where project_id = ?)", p.Id).Error
}

