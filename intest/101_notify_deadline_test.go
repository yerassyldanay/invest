package intest

import (
	"context"
	"fmt"
	"invest/model"
	"testing"
	"time"
)

func TestNotifyDeadline(t *testing.T) {
	cnx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
	model.OnlyNotifyAboutGantaDeadline(cnx)

	time.Sleep(time.Second * 100)
	defer cancel()
}

func TestNotifyDocumentDeadline(t *testing.T) {
	document := model.Document{}
	documents, err := document.OnlyGetEmptyDocumentsWithComingDeadline()
	switch {
	case err != nil:
		t.Error(err)
	case len(documents) < 1:
		fmt.Println("[WARN] there is no any deadline in documents")
		return
	}

	ndd := model.NotifyDocDeadline{
		DocumentId: documents[0].Id,
		//Document: documents[0],
	}

	model.GetMailerQueue().NotificationChannel <- &ndd

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 20)
	go model.GetMailerQueue().Handle(ctx)

	time.Sleep(time.Second * 50)
	defer cancel()
}
