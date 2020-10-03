package model

import (
	"invest/utils"
)


func (p *Project) Get_this_project_by_project_id() (utils.Msg) {
	err := p.GetAndUpdateStatusOfProject(GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	err = p.OnlyGetByIdPreloaded(GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	err = p.OnlyGetCategorsByProjectId(GetDB())
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	err = p.OnlyUnmarshalInfo()
	if err != nil {
		return ReturnInternalDbError(err.Error())
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return ReturnNoErrorWithResponseMessage(resp)
}
