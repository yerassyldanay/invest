package service

import (
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Assign(pu model.ProjectsUsers) (utils.Msg) {

	var trans = model.GetDB().Begin()
	defer func() { if trans != nil { trans.Rollback() } }()

	// check if there is a such project
	var project = model.Project{Id: pu.ProjectId}
	if err := project.OnlyGetById(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// check user
	var user = model.User{Id: pu.UserId}
	if err := user.OnlyGetUserById(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// create project & user relation
	pu.Created = utils.GetCurrentTime()
	if err := pu.OnlyCreate(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// set this step to done
	var ganta = model.Ganta{
		ProjectId: project.Id,
	}
	if err := ganta.OnlyChangeStatusToDoneById(trans); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	if err := trans.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}
