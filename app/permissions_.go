package app

import (
	"fmt"
	"invest/model"
	"time"
)

type per struct {
	str			string
}

func (p *per) p(err error) {
	if err != nil {
		fmt.Println(p.str, err.Error())
	}
}

func Prepare_permissions() {
	var p = per{str: "Prepare_permissions"}
	var all_perm = []string{
		Perm_1_crud_user,
		Perm_2_projects_see_all,
		Perm_3_projects_see_own,
		Perm_4_projects_make_changes,
		Perm_5_projects_comment,
		Perm_6_projects_accept,
		Perm_7_analysis_see,
		Perm_8_projects_submit,
	}

	for i, perm := range all_perm {
		p.p(model.GetDB().Create(&model.Permission{
			Id:          	uint64(i+1),
			Name:        	perm,
			Description: 	PMapPermissionMap[perm],
		}).Error)
		time.Sleep(time.Second * 0)
	}
}