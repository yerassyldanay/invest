package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
	"sync"
)

func (p *Project) OnlyGetProjectByUserId(user_id uint64, trans *gorm.DB) (projects []Project, err error) {
	err = trans.Raw("select p.* from projects p join projects_users pu on p.id = pu.project_id where user_id = ?;", user_id).
		Scan(&projects).Error
	return projects, err
}

func (p *Project) Get_projects_by_user_id(user_id uint64) (utils.Msg) {
	projects, err := p.OnlyGetProjectByUserId(user_id, GetDB())

	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var wg = sync.WaitGroup{}
	for i, _ := range projects {
		i := i
		wg.Add(1)

		go func (project *Project) {
			defer wg.Done()
			if err := project.OnlyGetCategorsByProjectId(GetDB()); err != nil {
				project.Categors = []Categor{}
			}
		}(&projects[i])

	}
	wg.Wait()

	var projectsMap = []map[string]interface{}{}
	for _, project := range projects {
		projectsMap = append(projectsMap, Struct_to_map(project))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = projectsMap

	return utils.Msg{resp, 200, "", ""}
}
