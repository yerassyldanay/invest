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
	var main_query = `select u.* from users u join roles r on u.role_id = r.id where r.name = 'admin';`
	_ = GetDB().Exec(main_query).Scan(&users).Error

	return users
}

