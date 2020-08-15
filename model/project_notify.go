package model

import (
	"invest/utils"
)

func (p *Project) Notify_all_assigned_users_and_admin(message SendgridMessage) (map[string]interface{}, error) {
	var main_query =
		` with user_structs as ( ` +
		` select u.fio as fio, e.address as address from projects_users pu join users u on pu.user_id = u.id ` +
		` join roles r on u.role_id = r.id ` +
		` join emails e on u.email_id = e.id ` +
		` where r.name != 'investor' and project_id = ? ` +
		` union ( ` +
		` select u.fio as fio, e.address as address from users u  ` +
		` join roles r2 on u.role_id = r2.id ` +
		` join emails e on u.email_id = e.id  ` +
		` where r2.name = 'admin' ` +
		` ) ` +
		` ) select * from user_structs; `

	type userStruct struct{
		Fio				string				`json:"fio"`
		Address			string				`json:"address"`
	}

	var check = []userStruct{}

	err := GetDB().Raw(main_query, p.Id).Scan(&check).Error
	if err != nil {
		return utils.ErrorInternalDbError, err
	}

	var emails, names = []string{}, []string{}

	for _, each := range check {
		emails = append(emails, each.Address)
		names = append(names, each.Fio)
	}

	var sms = SendgridMessageStore {
		ToList:            		emails,
		ToListName:        		names,
		SendgridMessage:   		message,
	}

	return sms.SendMessageToList()
}

