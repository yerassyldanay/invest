package model

import "invest/utils"

func (si *SignIn) Is_on_db() bool {
	if ok := utils.Is_it_free_from_sql_injection(si.KeyUsername); !ok {
		return false
	}

	var count int
	if err := GetDB().Table(User{}.TableName()).Where(si.KeyUsername + "=?", si.Value).Count(&count).Error;
		err != nil {
			return false
	}

	return count == 0
}
