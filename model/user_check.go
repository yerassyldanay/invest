package model

/*
	check whether a user has a following role
 */
func (u *User) Is_this_is_role_of_user(role string) (bool) {
	if err := GetDB().Preload("Role").Where("id = ?", u.Id).First(u).Error;
		err != nil {
			return false
	}

	return u.Role.Name == role
}
