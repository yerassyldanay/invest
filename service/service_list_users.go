package service

import (
	"invest/model"
	"invest/utils/errormsg"
	"invest/utils/message"
)

func (is *InvestService) Get_project_with_its_users(project_id uint64) (message.Msg) {
	var err error
	var project = model.Project{Id: project_id}

	if err = project.OnlyGetByIdPreloaded(model.GetDB()); err != nil {
		// error occurred
	} else if err = project.OnlyGetCategorsByProjectId(model.GetDB()); err != nil {
		// error occurred
	} else if err = project.OnlyGetAssignedUsersByProjectId(model.GetDB()); err != nil {
		// error occurred
	} else {
		// no err, which means the following:
		// we have info on the project & categories as well as all assigned users
		var resp = errormsg.NoErrorFineEverthingOk

		// get rid of password
		for i, _ := range project.Users {
			project.Users[i].Password = ""
		}

		resp["info"] = model.Struct_to_map(project)

		return model.ReturnNoErrorWithResponseMessage(resp)
	}

	return model.ReturnInternalDbError(err.Error())
}
