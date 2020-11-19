package tests

import (
	"fmt"
	"invest/model"
	"invest/service"
	"testing"
)

func TestServiceProjectGetAllAssignedUsersToProject(t *testing.T) {
	var project = getAnyProject(t)

	// headers
	is := service.InvestService{BasicInfo:service.BasicInfo{UserId: 1}} // admin

	// get the list of projects
	msg := is.Get_project_with_its_users(project.Id)
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}

	// get users
	if err := project.OnlyGetAssignedUsersByProjectId(model.GetDB()); err != nil {
		t.Error(err)
	} else if len(project.Users) < 1 {
		fmt.Println("[WARN] the number of assigned users is 0. project: ", project.Id)
	}
}

func TestServiceProjectGetListOfProjects(t *testing.T) {
	// prepare status
	statuses := model.Prepare_project_statuses("")

	// logic
	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId: 1,
		},
	}

	// get projects
	msg := is.Get_projects_by_user_id_and_status(4, statuses, []int{1, 2, 3, 4}) // expert
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
}

func TestServiceProjectGetOwnProjects(t *testing.T) {
	// statuses
	statuses := model.Prepare_project_statuses("")

	// get project
	project := model.Project{}
	projects, err := project.OnlyGetProjectsOfInvestor(3, statuses, []int{1, 2, 3, 4}, "0", model.GetDB())
	if err != nil {
		t.Error(err)
	} else if len(projects) < 1 {
		fmt.Println("[WARN] a number of projects is 0")
	} else {
		fmt.Println("found ", len(projects), " projects")
	}

	// get project of spk user
	projects, err = project.OnlyGetProjectsOfSpkUsers(2, statuses, []int{1, 2, 3, 4}, "0", model.GetDB())
	if err != nil {
		t.Error(err)
	} else if len(projects) < 1 {
		fmt.Println("[WARN] a number of projects is 0")
	} else {
		fmt.Println("found ", len(projects), " projects")
	}
}

// spk users can get all projects
func TestServiceGetAllProjects(t *testing.T) {
	is := service.InvestService{
		BasicInfo: service.BasicInfo{
			UserId: 1,
		},
	}

	statuses := model.Prepare_project_statuses("")

	// get all projects by admin
	msg := is.Get_all_projects_by_statuses(statuses, []int{1, 2, 3, 4})
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}

	// get all projects by manager
	is.BasicInfo.UserId = 2 // manager
	msg = is.Get_all_projects_by_statuses(statuses, []int{1, 2, 3, 4})
	if msg.IsThereAnError() {
		t.Error(msg.ErrMsg)
	}
}

func TestModelGetProjects(t *testing.T) {
	// status
	statuses := model.Prepare_project_statuses("")

	//
	var project = model.Project{}
	projects, err := project.OnlyGetProjectsByStatusesAndSteps("0", statuses, []int{1, 2, 3, 4}, model.GetDB())

	switch {
	case err != nil:
		t.Error(err)
	case len(projects) < 1:
		fmt.Println("[WARN] a manager could not get any project")
	default:
		fmt.Println("A manager got ", len(projects), " project(s)")
	}
}
