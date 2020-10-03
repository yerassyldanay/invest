package service

import (
	"github.com/jinzhu/gorm"
	"invest/model"
	"invest/utils"
	"sync"
)

func (is *InvestService) Get_own_projects(statuses []string) (utils.Msg) {
	var err error
	var project = model.Project{}
	var projects []model.Project

	switch {
	case is.RoleName == utils.RoleInvestor:
		projects, err = project.OnlyGetProjectsOfInvestor(is.UserId, statuses, is.Offset, model.GetDB())
	default:
		projects, err = project.OnlyGetProjectsOfSpkUsers(is.UserId, statuses, is.Offset, model.GetDB())
	}

	// if there are no projects then return empty list
	if err == gorm.ErrRecordNotFound {
		projects = []model.Project{}
	}

	// preload categories
	var wg = sync.WaitGroup{}
	for i, _ := range projects {
		i := i
		wg.Add(1)
		go func(proj *model.Project, wg *sync.WaitGroup) {
			defer wg.Done()
			_ = proj.OnlyGetCategorsByProjectId(model.GetDB())
		}(&projects[i], &wg)
	}
	wg.Wait()

	// convert them to map
	var projectsMap = []map[string]interface{}{}
	for _, project := range projects {
		_ = project.OnlyUnmarshalInfo()
		project.Info = ""
		projectsMap = append(projectsMap, model.Struct_to_map(project))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = projectsMap

	return model.ReturnNoErrorWithResponseMessage(resp)
}

func (is *InvestService) Get_projects_by_user_id_and_status(user_id uint64, statuses []string) (utils.Msg) {
	var user = model.User{Id: user_id}
	if err := user.OnlyGetByIdPreloaded(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// make this as a request of the user (with user_id), not admin
	is.UserId = user_id
	is.RoleName = user.Role.Name

	// reuse
	return is.Get_own_projects(statuses)
}

func (is *InvestService) Get_all_projects_by_statuses(statuses []string) (utils.Msg) {
	var project = model.Project{}

	// get projects
	projects, _ := project.OnlyGetProjectsByStatuses(is.Offset, statuses, model.GetDB())

	// get categories
	var wg = sync.WaitGroup{}
	var projectsMap = []map[string]interface{}{}
	for i, _ := range projects {
		wg.Add(1)
		go func(proj *model.Project, gwg *sync.WaitGroup) {
			defer gwg.Done()
			_ = proj.OnlyGetCategorsByProjectId(model.GetDB())
		}(&projects[i], &wg)
	}
	wg.Wait()

	// convert to map
	for _, project := range projects {
		projectsMap = append(projectsMap, model.Struct_to_map(project))
	}

	// convert
	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = projectsMap

	return model.ReturnNoErrorWithResponseMessage(resp)
}
