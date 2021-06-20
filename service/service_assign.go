package service

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/message"
)

func (is *InvestService) AssignUserToProject(pu model.ProjectsUsers) message.Msg {

	// tx
	tx := model.GetDB().Begin()

	// create relation
	if err := pu.OnlyCreate(tx); err != nil {
		_ = tx.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}

	// get user role
	user := model.User{Id: pu.UserId}
	if err := user.OnlyGetByIdPreloaded(tx); err != nil {
		_ = tx.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}

	// if this is a manager
	if user.Role.Name == constants.RoleManager {
		project := model.Project{
			Id:                pu.ProjectId,
			IsManagerAssigned: true,
		}

		if err := project.OnlyUpdateById(tx, "is_manager_assigned"); err != nil {
			_ = tx.Rollback()
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

func (is *InvestService) AssignRemoveRelation(pu model.ProjectsUsers) message.Msg {

	// tx
	tx := model.GetDB().Begin()

	// delete relation
	if err := tx.Delete(&model.ProjectsUsers{}, "project_id = ? and user_id = ?", pu.ProjectId, pu.UserId).
		Error; err != nil {
		_ = tx.Rollback()
		return model.ReturnInternalDbError(err.Error())
	}

	// count number of managers assigned to the project
	number, err := pu.OnlyCountByRoleAndProjectId(constants.RoleManager, tx)
	switch {
	case err != nil:
		_ = tx.Rollback()
		return model.ReturnInternalDbError(err.Error())
	case number == 0:
		// set is_manager_assigned to false
		if err := tx.Model(&model.Project{}).Where("id = ?", pu.ProjectId).
			Updates(map[string]interface{}{
				"is_manager_assigned": false,
			}).Error; err != nil {
			_ = tx.Rollback()
			return model.ReturnInternalDbError(err.Error())
		}
	}

	// commit
	if err := tx.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// ok
	return model.ReturnNoError()
}
