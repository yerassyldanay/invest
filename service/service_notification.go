package service

import (
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
)

func (is *InvestService) Notification_get_by_project_id(project_id uint64) (message.Msg) {
	// get user email address
	var user = model.User{Id: is.UserId}
	if err := user.OnlyGetByIdPreloaded(model.GetDB()); err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	ni := model.NotificationInstance{
		ToAddress:      user.Email.Address,
	}
	
	// get notifications
	notifications, err := ni.OnlyGetNotificationsByEmailAndProjectId(ni.ToAddress, project_id, is.Offset, model.GetDB())
	if err != nil {
		return model.ReturnInternalDbError(err.Error())
	}

	// convert all notifications to map
	notificationsMap := []map[string]interface{}{}
	for _, notification := range notifications {
		notificationsMap = append(notificationsMap, model.Struct_to_map(notification))
	}

	var resp = errormsg.NoErrorFineEverthingOk
	resp["info"] = notificationsMap

	return model.ReturnNoErrorWithResponseMessage(resp)
}
