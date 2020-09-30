package service

import (
	"fmt"
	"invest/model"
	"invest/utils"
	"os"
)

func (is *InvestService) Add_documents_to_project(document *model.Document) (utils.Msg) {
	/*
		you cannot add document if:
			* this is a parent ganta step
			* a ganta step already has a document
	*/
	var ganta = model.Ganta{Id: document.GantaId}
	_ = ganta.OnlyGetById(model.GetDB())

	// check whether these requirements are met
	if ganta.GantaParentId == 0 || ganta.Does_this_ganta_step_has_document(model.GetDB()) {
		return model.ReturnMethodNotAllowed("the ganta step already possesses a document")
	}

	// responsible = investor or spk (manager or expert)
	switch {
	case ganta.Responsible == utils.RoleSpk && is.RoleName != utils.RoleInvestor:
		// pass
	case ganta.Responsible == utils.RoleInvestor && is.RoleName == utils.RoleInvestor:
		// pass
	default:
		return model.ReturnMethodNotAllowed("responsible: " + ganta.Responsible + " | your role: " + is.RoleName)
	}

	// store on db
	msg := document.Add_after_validation()

	// remove document in case of an error
	if msg.ErrMsg != "" {
		if err := os.Remove("." + document.Uri); err != nil {
			fmt.Println(err.Error())
		}
	}

	msg = model.ReturnNoError()
	return msg
}
