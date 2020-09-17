package model

import (
	"invest/utils"
	"sync"
)

func (p *Project) Get_this_project_with_its_users() (utils.Msg) {
	if err := p.Get_by_id(GetDB()); err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	_ = p.Get_all_categors_by_project_id(GetDB())
	_ = p.Unmarshal_info()

	if err := p.Get_only_assigned_users_to_project(GetDB()); err != nil {
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
	err := p.Get_by_id(GetDB())
	err2 := p.Get_all_categors_by_project_id(GetDB())

	if err != nil || err2 != nil {
		var errmsg string

		if err != nil { errmsg += err.Error() + " " }
		if err != nil { errmsg += err.Error() }

		return utils.Msg{utils.ErrorInternalDbError, 417, "", errmsg}
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = Struct_to_map(*p)

	return utils.Msg{resp, 200, "", ""}
}
