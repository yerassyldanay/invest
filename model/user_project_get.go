package model

import (
	"encoding/json"
	"fmt"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
)

/*
	get all projects for admin only
*/
func (u *User) Get_all_projects(offset string) (map[string]interface{}, error) {
	rows, err := GetDB().Table(Project{}.TableName()).Offset(offset).Limit(GetLimit).Rows()
	if err != nil {
		return errormsg.ErrorInternalDbError, err
	}
	defer rows.Close()

	type Tstruct struct {
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

		var main_query = "select u.fio, u.position from projects_users pu " +
			" join users u on pu.user_id = u.id " +
			" join roles r on u.role_id = r.id where pu.project_id=?; "
		err = GetDB().Exec(main_query, tproject.Id).Omit("password").Find(&tproject.Users).Error

		for i, _ := range tproject.Users {
			tproject.Users[i].Password = ""
		}

		//fmt.Println("Get projects: ", err)

		projects.Info = append(projects.Info, tproject)
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(projects)

	return resp, nil
}

/*
	get only own projects
*/
func (u *User) Get_own_projects_spk(offset string) (message.Msg) {
	var result = []map[string]interface{}{}
	var projects = []Project{}

	if err := u.OnlyGetByIdPreloaded(GetDB()); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	/*
		if this is not an investor
	*/
	var main_query = "select p.* from projects p join projects_users pu on p.id = pu.project_id where pu.user_id=? " +
		" offset ? limit ?;"

	err := GetDB().Raw(main_query, u.Id, offset, GetLimit).Find(&projects).Error
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	fmt.Println("Num of projects assigned to this user (if it is): ", len(projects))

	for _, project := range projects {
		_ = project.OnlyGetCategorsByProjectId(GetDB())
		_ = json.Unmarshal([]byte(project.Info), &project.InfoSent)
		result = append(result, Struct_to_map_with_escape(project, []string{"documents"}))
	}

	/*
		if this is an investor
	*/
	projects = []Project{}
	_ = GetDB().Preload("Organization").Find(&projects, "offered_by_id = ?", u.Id).Offset(offset).Limit(GetLimit).Error

	//fmt.Println("Num of projects that are offered by this investor (if it is): ", len(projects))

	for _, project := range projects {
		_ = project.OnlyGetCategorsByProjectId(GetDB())
		_ = json.Unmarshal([]byte(project.Info), &project.InfoSent)
		result = append(result, Struct_to_map_with_escape(project, []string{"documents"}))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = result

	return ReturnNoError()
}

/*
	get only own projects
*/
func (u *User) Get_own_projects_spk_(offset string) (message.Msg) {
	var result = []map[string]interface{}{}
	var projects = []Project{}

	if err := u.OnlyGetByIdPreloaded(GetDB()); err != nil {
		return ReturnInternalDbError(err.Error())
	}

	/*
		if this is not an investor
	*/
	var main_query = "select p.* from projects p join projects_users pu on p.id = pu.project_id where pu.user_id=? " +
		" offset ? limit ?;"

	err := GetDB().Raw(main_query, u.Id, offset, GetLimit).Find(&projects).Error
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	fmt.Println("Num of projects assigned to this user (if it is): ", len(projects))

	for _, project := range projects {
		_ = project.OnlyGetCategorsByProjectId(GetDB())
		_ = json.Unmarshal([]byte(project.Info), &project.InfoSent)
		result = append(result, Struct_to_map_with_escape(project, []string{"documents"}))
	}

	/*
		if this is an investor
	*/
	projects = []Project{}
	_ = GetDB().Preload("Organization").Find(&projects, "offered_by_id = ?", u.Id).Offset(offset).Limit(GetLimit).Error

	//fmt.Println("Num of projects that are offered by this investor (if it is): ", len(projects))

	for _, project := range projects {
		_ = project.OnlyGetCategorsByProjectId(GetDB())
		_ = json.Unmarshal([]byte(project.Info), &project.InfoSent)
		result = append(result, Struct_to_map_with_escape(project, []string{"documents"}))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = result

	return ReturnNoError()
}

/*
	preload & get all projects
 */
func (p *Project) Get_all_after_preload(offset string) (message.Msg) {
	var projects = []Project{}
	if err := GetDB().Preload("Organization").Preload("Categors").
		Offset(offset).Limit(constants.GetLimitProjects).Find(&projects).Error;
		err != nil {
			return message.Msg{errormsg.ErrorInternalDbError, 417, "", err.Error()}
	}

	var rprojects = []map[string]interface{}{}
	for _, project := range projects {
		_ = json.Unmarshal([]byte(project.Info), &project.InfoSent)
		project.Info = ""
		rprojects = append(rprojects, Struct_to_map(project))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = rprojects

	return message.Msg{resp, 200, "", ""}
}
