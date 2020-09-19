package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

func (p *Project) Only_get_projects_by_user_id(user_id uint64, trans *gorm.DB) (projects []Project, err error) {
	err = trans.Raw("select p.* from projects p join projects_users pu on p.id = pu.project_id where user_id = ?;", user_id).
		Scan(&projects).Error
	return projects, err
}

func (p *Project) Get_projects_by_user_id(user_id uint64) (utils.Msg) {
	projects, err := p.Only_get_projects_by_user_id(user_id, GetDB())

	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var projectsMap = []map[string]interface{}{}
	for _, project := range projects {
		projectsMap = append(projectsMap, Struct_to_map(project))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = projectsMap

	return utils.Msg{resp, 200, "", ""}
}
