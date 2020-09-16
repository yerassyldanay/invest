package model

import (
	"encoding/json"
	"invest/utils"
)

/*
	get all projects for admin only
*/
func (u *User) Get_all_projects(offset string) (map[string]interface{}, error) {
	rows, err := GetDB().Table(Project{}.TableName()).Offset(offset).Limit(GetLimit).Rows()
	if err != nil {
		return utils.ErrorInternalDbError, err
	}
	defer rows.Close()

	type Tstruct struct {
		Username			string
		Position			string
		Fio					string
		Role				string
	}

	var projects = struct{
		Info			[]Project
	}{}
	var tproject = Project{}
	for rows.Next() {
		if err := GetDB().ScanRows(rows, &tproject); err != nil {
			continue
		}

		var main_query = "select u.username, u.fio, u.position from projects_users pu " +
			" join users u on pu.user_id = u.id " +
			" join roles r on u.role_id = r.id where pu.project_id=?; "
		err = GetDB().Exec(main_query, tproject.Id).Omit("password").Find(&tproject.Users).Error

		for i, _ := range tproject.Users {
			tproject.Users[i].Password = ""
		}

		//fmt.Println("Get projects: ", err)

		projects.Info = append(projects.Info, tproject)
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(projects)

	return resp, nil
}

/*
	get only own projects
*/
func (u *User) Get_own_projects(offset string) (map[string]interface{}, error) {
	/*
		if this is not an investor
	 */
	var main_query = "select p.* from projects p join projects_users pu on p.id = pu.project_id where pu.user_id=? " +
		" offset ? limit ?;"

	var projects = []Project{}
	err := GetDB().Exec(main_query, u.Id, offset, GetLimit).Find(&projects).Error
	if err != nil {
		return utils.ErrorInternalDbError, err
	}

	var result = []map[string]interface{}{}
	for _, project := range projects {
		_ = json.Unmarshal([]byte(project.Info), &project.InfoSent)
		result = append(result, Struct_to_map(project))
	}

	/*
		if this is an investor
	 */
	projects = []Project{}
	_ = GetDB().Find(&projects, "offered_by_id = ?", u.Id).Error

	for _, project := range projects {
		_ = json.Unmarshal([]byte(project.Info), &project.InfoSent)
		result = append(result, Struct_to_map(project))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = result

	return resp, nil
}

/*
	preload & get all projects
 */
func (p *Project) Get_all_after_preload(offset string) (utils.Msg) {
	var projects = []Project{}
	if err := GetDB().Preload("Organization").Preload("Categors").
		Offset(offset).Limit(utils.GetLimitProjects).Find(&projects).Error;
		err != nil {
			return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var rprojects = []map[string]interface{}{}
	for _, project := range projects {
		rprojects = append(rprojects, Struct_to_map(project))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = rprojects

	return utils.Msg{resp, 200, "", ""}
}
