package model

import (
	"invest/utils"
	"sync"
)

func (p *Project) Get_this_project_with_its_users() (utils.Msg) {
	if err := p.OnlyGetById(GetDB()); err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	_ = p.OnlyGetCategorsByProjectId(GetDB())
	_ = p.OnlyUnmarshalInfo()

	if err := p.OnlyGetAssignedUsers(GetDB()); err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "",err.Error()}
	}

	var wg = sync.WaitGroup{}
	for i, _ := range p.Users {
		wg.Add(1)
		p.Users[i].Password = ""
		go p.Users[i].Add_statistics_to_this_user_on_project_statuses(&wg)
	}
	wg.Wait()

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return utils.Msg{resp, 200, "", ""}
}

func (p *Project) Get_this_project_by_project_id() (utils.Msg) {
	err := p.OnlyGetById(GetDB())
	if err != nil { return ReturnInternalDbError(err.Error())}

	err = p.OnlyGetCategorsByProjectId(GetDB())
	if err != nil { return ReturnInternalDbError(err.Error())}

	err = p.OnlyUnmarshalInfo()
	if err != nil { return ReturnInternalDbError(err.Error()) }

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return ReturnNoErrorWithResponseMessage(resp)
}
