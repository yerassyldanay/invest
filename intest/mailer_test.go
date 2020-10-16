package intest

import (
	"context"
	"invest/model"
	"invest/service"
	"testing"
	"time"
)

/*
	TestMailerQueue:
		this function tests whether the mailer queue is working properly
	Mailer Queue is a queue, which accepts notifications of various structure (but follows InterMessage)
		& handle all further job (create notification on db, prepare smtp message & send that message to an email address)
 */
func TestMailerQueue(t *testing.T) {
	notification := model.NotifyCreateProfile{
		UserId:      2,
		CreatedById: 1,
	}

	mq := model.GetMailerQueue()
	mq.NotificationChannel <- &notification

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	mq.Handle(ctx)

	time.Sleep(time.Second * 5)

	defer cancel()
}

func TestModelGetNotification(t *testing.T) {
	ni := model.NotificationInstance{
		ToAddress:      "invest.dept.spk@inbox.ru",
		Notification:   model.Notification{
			ProjectId: 1,
		},
	}

	notis, err := ni.OnlyGetNotificationsByEmailAndProjectId(ni.ToAddress, ni.Notification.ProjectId, model.GetDB())

	switch {
	case err != nil:
		t.Error("expected no error, but got ", err)
	case len(notis) < 1:
		t.Error("expected some notifications (at least one), but got 0")
	default:
		//fmt.Println(notis)
	}
}

func TestServiceGetNotifications(t *testing.T) {
	is := service.InvestService{BasicInfo:service.BasicInfo{UserId: 2}}

	// test
	msg := is.Notification_get_by_project_id(0)
	if msg.IsThereAnError() {
		t.Error("expected no error, but got ", msg.ErrMsg)
	}
}
