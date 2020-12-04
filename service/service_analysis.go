package service

import (
	"invest/model"
	"invest/utils/errormsg"
	"invest/utils/helper"
	"invest/utils/message"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// copy map
func OnlyCopySpecificMap(m map[string]int64) (map[string]int64) {
	newMap := map[string]int64{}
	for key, value := range m {
		newMap[key] = value
	}

	return newMap
}

// analysis
func (is *InvestService) Analysis_get_on_projects(analysis model.Analysis) (message.Msg) {
	// convert timestamp to date
	switch {
	case analysis.Start < 1:
		// then take for 5 years
		analysis.StartDate = helper.GetCurrentTime().Add(time.Hour * 24 * (-365) * 5)
	default:
		// convert it to
		analysis.StartDate = helper.OnlyPrettifyTime(time.Unix(analysis.Start, 0))
	}

	// convert end date
	switch {
	case analysis.End < 1:
		// then set current time
		analysis.EndDate = helper.GetCurrentTime()
	default:
		// convert
		analysis.EndDate = time.Unix(analysis.End, 0)
	}

	// get projects
	projects, err := analysis.OnlyGetProjectsByExtendedFields(model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

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
		analysis.ProjectExtendedList = append(analysis.ProjectExtendedList, projExtended)
	}

	var wg = sync.WaitGroup{}

	// load cost & finance tables
	for i, _ := range analysis.ProjectExtendedList {
		i := i
		wg.Add(1)
		go func(proj *model.ProjectExtended, gwg *sync.WaitGroup) {
			defer gwg.Done()
			_ = proj.Cost.OnlyGetByProjectId(model.GetDB())
			_ = proj.Finance.OnlyGetByProjectId(model.GetDB())
			_ = proj.OnlyGetCategorsByProjectId(model.GetDB())
			_ = proj.OnlyPreloadOrganizationByOrganizationId(model.GetDB())

			// check categories
			if len(proj.Categors) < 1 {
				proj.Categors = []model.Categor{}
			}

		}(&analysis.ProjectExtendedList[i], &wg)
	}
	wg.Wait()

	// return response
	var resp = errormsg.NoErrorFineEverthingOk
	if analysis.WriteToFile {

		// create file name + indicate file pather
		fileName := helper.Generate_Random_String(40)
		filePath, err := filepath.Abs("./documents/analysis/" + fileName + ".xlsx")
		if err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// write to this file
		err = analysis.OnlyWriteDataToFile(filePath, is.Lang)
		if err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// indicate in in info
		resp["info"] = "/documents/analysis/" + fileName + ".xlsx"

	} else {
		// get categories
		c := model.Categor{}
		categories, err := c.OnlyGetAll(model.GetDB())
		if err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		// create map
		categoriesMap := map[string]int64{}
		for _, c = range categories {
			key := strings.ToLower(c.Eng)
			categoriesMap[key] = 0
		}

		// pie-chart data
		pie_chart := map[string]map[string]int64{
			"employee_count": OnlyCopySpecificMap(categoriesMap),
			"tax": OnlyCopySpecificMap(categoriesMap),
			"investment": OnlyCopySpecificMap(categoriesMap),
			"count": OnlyCopySpecificMap(categoriesMap),
		}

		// convert it to map
		var projectsMapList = []map[string]interface{}{}
		for _, proj := range analysis.ProjectExtendedList {
			// convert to map
			projectsMapList = append(projectsMapList, model.Struct_to_map(proj))

			// prepare three pie charts
			for i, _ := range proj.Categors {
				category := strings.ToLower(proj.Categors[i].Eng)

				pie_chart["employee_count"][category] += int64(proj.EmployeeCount)
				pie_chart["tax"][category] += int64(proj.Finance.Taxes)
				pie_chart["investment"][category] += int64(proj.Cost.WorkingCapitalInvolved + proj.Cost.WorkingCapitalInvestor)
				pie_chart["count"][category] += 1
			}

			//fmt.Println(proj.Finance.Taxes, proj.Cost.WorkingCapitalInvolved + proj.Cost.WorkingCapitalInvestor, proj.EmployeeCount)
		}

		// enter to the response map
		resp["info"] = projectsMapList
		resp["pie_chart"] = pie_chart
	}

	return model.ReturnNoErrorWithResponseMessage(resp)
}
