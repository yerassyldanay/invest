package service

import (
	"invest/model"
	"invest/utils"
)

func (is *InvestService) Check_whether_this_user_is_responsible_for_document(document_id uint64, project_id uint64) (utils.Msg) {
	var document = model.Document{Id: document_id, ProjectId: project_id}

	// get this document
	err := document.OnlyGetDocumentById(model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// convert role to spk
	switch is.RoleName {
	case utils.RoleManager:
		is.RoleName = utils.RoleSpk
	case utils.RoleExpert:
		is.RoleName = utils.RoleSpk
	}

	// check whether this user is responsible
	if document.Responsible != is.RoleName {
		return model.ReturnMethodNotAllowed("responsible: " + document.Responsible + " | your role: " + is.RoleName)
	}

	return model.ReturnNoError()
}

