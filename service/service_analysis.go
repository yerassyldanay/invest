package service

import (
	"invest/model"
	"invest/utils"
	"sync"
	"time"
)

// analysis
func (is *InvestService) Analysis_get_on_projects(analysis model.Analysis) (utils.Msg) {
	// convert timestamp to date
	switch {
	case analysis.Start < 1:
		// then take for 5 years
		analysis.StartDate = utils.GetCurrentTime().Add(time.Hour * 24 * (-365) * 5)
	default:
		// convert it to
		analysis.StartDate = time.Unix(analysis.Start, 0)
	}

	// convert end date
	switch {
	case analysis.End < 1:
		// then set current time
		analysis.EndDate = utils.GetCurrentTime()
	default:
		// convert
		analysis.EndDate = time.Unix(analysis.End, 0)
	}

	// get projects
	projects, err := analysis.Get_projects_by_steps(is.Offset, model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	var projectsList = []model.ProjectExtended{}
	var projExtended = model.ProjectExtended{}

	for _, project := range projects {
		projExtended = model.ProjectExtended{
			Project: project,
			Cost:    model.Cost{
				ProjectId: project.Id,
			},
			Finance: model.Finance{
				ProjectId: project.Id,
			},
		}
		projectsList = append(projectsList, projExtended)
	}

	var wg = sync.WaitGroup{}

	// load cost & finance tables
	for i, _ := range projectsList {
		i := i
		wg.Add(1)
		go func(proj *model.ProjectExtended, gwg *sync.WaitGroup) {
			defer gwg.Done()
			_ = proj.Cost.OnlyGetByProjectId(model.GetDB())
			_ = proj.Finance.OnlyGetByProjectId(model.GetDB())
		}(&projectsList[i], &wg)
	}
	wg.Wait()

	// convert it to map
	var projectsMapList = []map[string]interface{}{}
	for _, proj := range projectsList {
		projectsMapList = append(projectsMapList, model.Struct_to_map(proj))
	}

	// return response
	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = projectsMapList

	return model.ReturnNoErrorWithResponseMessage(resp)
}
