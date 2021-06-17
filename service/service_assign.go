package service

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/message"
)

func (is *InvestService) Assign_user_to_project(pu model.ProjectsUsers) (message.Msg) {

	tx := model.GetDB().Begin()
	defer func() { if tx != nil { tx.Rollback() } }()

	// create relation
	if err := pu.OnlyCreate(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// get user role
	user := model.User{Id: pu.UserId}
	if err := user.OnlyGetByIdPreloaded(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// if this is a manager
	if user.Role.Name == constants.RoleManager {
		project := model.Project{
			Id: pu.ProjectId,
			IsManagerAssigned: true,
		}

		if err := project.OnlyUpdateById(tx, "is_manager_assigned"); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}
	}

	// commit changes
	if err := tx.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// send message
	na := model.NotifyAssign{
		UserId:    pu.UserId,
		ProjectId: pu.ProjectId,
	}

	// this handles everything
	select {
	case model.GetMailerQueue().NotificationChannel <- &na:
	default:
	}

	return model.ReturnNoError()
}

func (is *InvestService) Assign_remove_relation(pu model.ProjectsUsers) (message.Msg) {

	// transaction cannot be used
	// as we will count until the result is committed
	if err := pu.OnlyDelete(model.GetDB()); err  != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// count number of managers assigned to the project
	number, err := pu.OnlyCountByRoleAndProjectId(constants.RoleManager, model.GetDB())
	switch {
	case err != nil:
		return model.ReturnInternalDbError(err.Error())
	case number == 0:
		project := model.Project{
			Id: pu.ProjectId,
		}

		// get project
		if err := project.OnlyGetById(model.GetDB()); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}

		project.IsManagerAssigned = false

		// save changes
		if err := project.OnlySave(model.GetDB()); err != nil {
			return model.ReturnInternalDbError(err.Error())
		}
	}

	return model.ReturnNoError()
}
