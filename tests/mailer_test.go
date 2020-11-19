package tests

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

	notis, err := ni.OnlyGetNotificationsByEmailAndProjectId(ni.ToAddress, ni.Notification.ProjectId, "0", model.GetDB())

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

func TestNotifyCreateProfile(t *testing.T) {
	nq := model.NotifyCreateProfile{
		UserId:      2,
		CreatedById: 1,
		RawPassword: "6z24HXMd7nLeZAE",
	}

	mq := model.GetMailerQueue()
	mq.NotificationChannel <- &nq

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	mq.Handle(ctx)

	time.Sleep(time.Second * 10)
	defer cancel()
}

func TestNotifyNewPassword(t *testing.T) {
	nnp := model.NotifyNewPassword{
		UserId:         2,
		RawNewPassword: "6z24HXMd7nLeZAE",
	}

	mq := model.GetMailerQueue()
	mq.NotificationChannel <- &nnp

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	mq.Handle(ctx)

	time.Sleep(time.Second * 10)
	defer cancel()
}