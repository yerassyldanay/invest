package intest

import (
	"context"
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
