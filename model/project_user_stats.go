package model

import "invest/utils"

/*
	1. передать статистику определенного сотрудника
	2. передать проекты по статусу
	3. передать проекты по статусу + по айди сотрудника
 */
func (pus *ProjectUserStat) Get_projects_by_status(offset string) (utils.Msg) {
	if offset == "" {offset = "0"}

	var projects = []Project{}
	if err := GetDB().Preload("Organization").Table(Project{}.TableName()).
		Where("status = ?", pus.Status).Offset(offset).Limit(GetLimit).Find(&projects).Error;
		err != nil {
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

/*
	respond with projects that correspond a certain status & user id
 */
func (pus *ProjectUserStat) Get_projects_by_status_and_user_id(offset string) (utils.Msg) {
	var main_query = `
		select p.* from projects p join projects_users pu
			on p.id = pu.project_id where pu.user_id = ?
			and p.status = ? offset ? limit ? ;
	`

	var projects = []Project{}
	err := GetDB().Raw(main_query, pus.UserId, pus.Status, offset, GetLimit).Scan(&projects).Error
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

