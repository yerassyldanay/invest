package service

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/helper"
	"testing"
)

func TestInvestService_AssignRemoveRelation(t *testing.T) {
	// create project
	p1 := helperCreateManagerWithProject(t)
	_ = p1
	p2 := helperCreateManagerWithProject(t)
	_ = p2

	// remove
	is := InvestService{}
	msg := is.AssignRemoveRelation(model.ProjectsUsers{
		ProjectId: p1.Project.Id,
		UserId:    p1.Manager.Id,
	})
	helper.HelperPrint(msg)

	// check
	var pu = model.ProjectsUsers{}
	_ = TestGorm.Model(&pu).Where("project_id = ? and user_id = ?", p1.Project.Id, p1.Manager.Id).
		First(&pu).Error
	require.Zero(t, pu.UserId)
	require.Zero(t, pu.ProjectId)

	// get second one
	// check
	err := TestGorm.First(&pu, "project_id = ? and user_id = ?", p2.Project.Id, p2.Manager.Id).Error
	require.NoError(t, err)
	require.Equal(t, pu.UserId, p2.Manager.Id)
	require.Equal(t, pu.ProjectId, p2.Project.Id)
}

func TestInvestService_AssignUserToProject(t *testing.T) {
	//
	projectElement := helperCreateManagerWithProject(t)

	// error
	is := InvestService{}
	msg := is.AssignUserToProject(model.ProjectsUsers{
		ProjectId: projectElement.Project.Id,
		UserId:    projectElement.Manager.Id,
	})
	require.NotZero(t, msg.ErrMsg)
	//helper.HelperPrint(msg)

	// get number of managers of a project
	pu := model.ProjectsUsers{}
	count, err := pu.OnlyCountByRoleAndProjectId(constants.RoleManager, TestGorm)
	require.NoError(t, err)

	// create a new manager & assign
	manager := helperTestCreateUser(constants.RoleManager, t)
	msg = is.AssignUserToProject(model.ProjectsUsers{
		ProjectId: projectElement.Project.Id,
		UserId:    manager.Id,
	})
	require.Zero(t, msg.ErrMsg)
	//helper.HelperPrint(msg)

	// get user
	require.NoError(t, TestGorm.First(&pu, "project_id = ? and user_id = ?",
		projectElement.Project.Id, manager.Id).Error)

	// get number of managers of a project
	newCount, err := pu.OnlyCountByRoleAndProjectId(constants.RoleManager, TestGorm)
	require.NoError(t, err)
	require.NotEqual(t, newCount, count)
}
