package model

/*
	get user info by admin
 */
func (c *User) Get_user_by_id() error {
	return GetDB().Table(c.TableName()).Where("id = ?", c.Id).First(c).Error
}

/*
	get admins
 */
func (c *User) Get_admins_only_user_info() (users []User){
	err := GetDB().Preload("Email").Preload("Phone").Table("users").
		Joins(" join roles on roles.id = users.role_id ").
		Where(" roles.name = 'admin' ").Find(&users).Error
	if err != nil {
		return users
	}

	return users
}

