package service

import (
	"invest/model"
	"invest/utils/errormsg"
	"invest/utils/message"
)

func (is *InvestService) SmtpCreate(smtp *model.SmtpServer) (message.Msg) {

	// validate
	if err := smtp.Validate(); err != nil {
		return model.ReturnInvalidParameters(err.Error())
	}

	// create transaction
	tx := model.GetDB().Begin()
	defer func() { if tx != nil { tx.Rollback() } }()

	// create smtp
	if err := smtp.OnlyCreate(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// create headers
	if err := smtp.OnlySetHeaders(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := tx.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}

// update
func (is *InvestService) SmtpUpdate(smtp *model.SmtpServer) (message.Msg) {

	// validate
	if err := smtp.Validate(); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// transaction
	tx := model.GetDB().Begin()
	defer func() { if tx != nil { tx.Rollback() } }()

	// delete all headers from db
	if err := smtp.OnlyDeleteAllHeadersBySmtpId(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// save new values
	if err := smtp.OnlySetHeaders(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// update values of smtp server
	if err := smtp.OnlySaveById(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := tx.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// no error
	return model.ReturnNoError()
}

// get smtp credentials
func (is *InvestService) SmtpGet() (message.Msg) {
	// get all smtp credentials
	smtp := model.SmtpServer{}
	smtps, err := smtp.OnlyGetAll(model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// convert to map
	smtpMap := []map[string]interface{}{}
	for _, smtp = range smtps {
		smtpMap = append(smtpMap, model.Struct_to_map(smtp))
	}

	resp := errormsg.NoErrorFineEverthingOk
	resp["info"] = smtpMap

	return model.ReturnNoErrorWithResponseMessage(resp)
}

// delete
func (is *InvestService) SmtpDelete(smtp *model.SmtpServer) (message.Msg) {

	// transaction
	tx := model.GetDB().Begin()
	defer func() { if tx != nil { tx.Rollback() } }()

	// delete headers
	if err := smtp.OnlyDeleteAllHeadersBySmtpId(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// delete smtp
	if err := smtp.OnlyDeleteById(tx); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// commit changes
	if err := tx.Commit().Error; err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	return model.ReturnNoError()
}
