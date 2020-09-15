package model

import "invest/utils"

func (p *Project) Get_this_project_with_its_users() (utils.Msg) {
	if err := GetDB().First(p, "id = ?", p.Id).Error; err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", ""}
	}

	if err := GetDB().Preload("Email").Omit("password, created").Preload("Role").Find(&p.Users, "id in (select user_id from projects_users where project_id = ?)", p.Id).Error;
		err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "",""}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return utils.Msg{resp, 200, "", ""}
}
