package model

import "invest/utils"

func (sms *SendgridMessageStore) Only_get_messages_by_project_id(offset interface{}) (sendgridMessageStores []SendgridMessageStore, err error) {
	err = GetDB().Preload("SendgridMessage").Find(&sendgridMessageStores, "project_id = ? and \"to\" = ?", sms.ProjectId, sms.To).Offset(offset).Error
	return sendgridMessageStores, err
}

func (sms *SendgridMessageStore) Get_messages_by_project_id(offset interface{}) (utils.Msg) {
	sendgridMessageStores, err := sms.Only_get_messages_by_project_id(offset)
	if err != nil {
		return utils.Msg{utils.ErrorInternalDbError, 417, "", err.Error()}
	}

	var smsMap = []map[string]interface{}{}
	for _, sms := range sendgridMessageStores {
		smsMap = append(smsMap, Struct_to_map(sms))
	}

	var resp = utils.NoErrorFineEverthingOk
	resp["info"] = smsMap

	return utils.Msg{resp, 200, "", ""}
}

