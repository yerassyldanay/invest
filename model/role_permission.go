package model

/*
	get permissions by id of role table
 */
func (r *Role) Get_permissions_by_role_id() error {
	return GetDB().Exec("select p.* from roles r inner join roles_permissions rp on r.id = rp.role_id " +
		"inner join permissions p on p.id = rp.permission_id where r.id =?", r.Id).Find(&r.Permissions).Error
}


