package tests

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"invest/model"
	"invest/service"
	"testing"
)

func TestServiceProjectGetAllAssignedUsersToProject(t *testing.T) {
	var project = HelperGetAnyProject(t)

	// headers
	is := service.InvestService{BasicInfo:service.BasicInfo{UserId: 1}} // admin

	// get the list of projects
	msg := is.Get_project_with_its_users(project.Id)

	// check
	require.Zero(t, msg.ErrMsg)

	// get users
	err := project.OnlyGetAssignedUsersByProjectId(model.GetDB())
	require.NoError(t, err)
	require.NotZero(t, len(project.Users))
	require.Condition(t, func() (bool) { return len(project.Users) > 0 })
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

	// check
	require.Zero(t, msg.ErrMsg)
}

func TestServiceProjectGetOwnProjects(t *testing.T) {
	// statuses
	statuses := model.Prepare_project_statuses("")

	// get project
	project := model.Project{}
	projects, err := project.OnlyGetProjectsOfInvestor(3, statuses, []int{1, 2, 3, 4}, "0", model.GetDB())

	// check
	require.NoError(t, err)
	require.Condition(t, func() (bool) { return len(projects) > 0 })

	// get project of spk user
	projects, err = project.OnlyGetProjectsOfSpkUsers(2, statuses, []int{1, 2, 3, 4}, "0", model.GetDB())

	// check
	require.NoError(t, err)
	if len(projects) < 1 {
		fmt.Println(">>> WARN  a number of projects is 0")
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

	// check
	require.Zero(t, msg.ErrMsg)

	// get all projects by manager
	is.BasicInfo.UserId = 2 // manager
	msg = is.Get_all_projects_by_statuses(statuses, []int{1, 2, 3, 4})

	// check
	require.Zero(t, msg.ErrMsg)
}

func TestModelGetProjects(t *testing.T) {
	// status
	statuses := model.Prepare_project_statuses("")

	//
	var project = model.Project{}
	projects, err := project.OnlyGetProjectsByStatusesAndSteps("0", statuses, []int{1, 2, 3, 4}, model.GetDB())

	// check
	require.NoError(t, err)
	if len(projects) < 1{
		HelperWarn("  a manager could not get any project")
	}
}
