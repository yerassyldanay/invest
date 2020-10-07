package service

import (
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Assign_user_to_project(pu model.ProjectsUsers) (utils.Msg) {

	// create relation
	if err := pu.OnlyCreate(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}

func (is *InvestService) Assign_remove_relation(pu model.ProjectsUsers) (utils.Msg) {
	if err := pu.OnlyDelete(model.GetDB()); err  != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}
