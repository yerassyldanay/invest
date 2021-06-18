package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/yerassyldanay/invest/model"
	"github.com/yerassyldanay/invest/utils/config"
	"github.com/yerassyldanay/invest/utils/helper"
	"testing"
	"time"
)

var TestRedis *redis.Client
var TestGorm *gorm.DB

func TestMain(m *testing.M) {
	var err error

	// environmental variables
	opts, err := config.LoadConfig("../environment/")
	helper.IfErrorPanic(err)
	_ = opts

	// POSTGRES - set up a connection with database
	TestGorm, err = model.EstablishDatabaseConnection(opts)
	helper.IfErrorPanic(err)

	// REDIS - establish connection with redis
	TestRedis, err = model.EstablishConnectionWithRedis(opts)
	helper.IfErrorPanic(err)
	TestRedis.FlushAll()

	// close the connection with database at the end
	defer func() {
		fmt.Println("[DATABASE] closing connection...")
		err = model.GetDB().Close()
		helper.IfErrorPanic(err)
	}()

	// setup mailer queue, which receives & handles notifications in one place
	ctx, cancel := context.WithCancel(context.Background())
	mq := model.InitiateNewMailerQueue()
	go mq.Handle(ctx)
	defer cancel()

	/*
		this will be main at the end of all
			allows to gracefully shut down
	*/
	defer time.Sleep(time.Millisecond * 10)

	// run notification sender at background
	cnx, cancelNotifier := context.WithCancel(context.Background())
	go model.OnlyNotifyAboutGantaDeadline(cnx)
	defer cancelNotifier()

	fmt.Println("[MAIN] starting tests...")

	m.Run()

	fmt.Println("[MAIN] ending tests...")
}
