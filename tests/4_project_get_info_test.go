package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/service"
	"testing"
)

//func

func TestProjectGetById(t *testing.T) {
	var newProject = HelperGetAnyProject(t)

	// check model
	project := model.Project{
		Id: newProject.Id,
	}
	err := project.Get_this_project_by_project_id()

	// check
	require.NoError(t, err)
	//require.NotZero(t, project.Organization.Id)
	require.NotZero(t, project.CurrentStep.Id)
	require.Condition(t, func() bool { return len(project.Users) > 0 })

	// go through users & check
	for _, user := range project.Users {
		require.NotZero(t, user.Fio)
		require.NotZero(t, user.Email.Address)
		require.NotZero(t, user.Phone.Number)
		require.NotZero(t, user.Phone.Ccode)
	}

	// service
	is := service.InvestService{BasicInfo: service.BasicInfo{UserId: 1}}
	msg := is.ProjectGetById(project.Id)

	// check service
	if msg.IsThereAnError() {
		t.Error(HelperExpectedNoErrorButGot(err.Error()))
	}
}
