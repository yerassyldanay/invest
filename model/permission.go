package model

import "github.com/jinzhu/gorm"

func (p *Permission) Check_permission_by_role_id(role_id uint64, perm string, tx *gorm.DB) (err error) {
	err = tx.Raw("select p.* from roles_permissions rp join permissions p on p.id = rp.permission_id where role_id = ? and p.name = ?", role_id, perm).
		Scan(p).Error
	return err
}
