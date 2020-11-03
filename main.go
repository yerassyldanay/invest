package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	logr "github.com/sirupsen/logrus"
	"invest/app"
	"invest/model"
	"invest/utils/constants"
	"invest/utils/logist"

	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	/*
		Set up a connection with db
	 */
	model.Set_up_db()

	/*
		setup mailer queue, which receives & handles notifications in one place
	 */
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mq := model.InitiateNewMailerQueue()
	go mq.Handle(ctx)

	/*
		migration
	*/
	if err := model.Migration(); err != nil {
		logist.SysMessage{
			FuncName: "MAIN",
			Message:  "migration: " + err.Error(),
			Ok:       false,
			Lev:      constants.WarnLevel,
		}.Log_system_message()
	}

	/*
		this will be run at the end of all
			allows to gracefully shut down
	 */
	defer time.Sleep(time.Millisecond * 10)

	/*
		close the file at the end of the
	*/
	logist.InitiateLogFile()
	defer logist.Get_file()

	/*
		this function stops the goroutines in the background
	 */
	defer func() {
		logist.Get_file_rotator().Cancel <- true
	}()

	/*
		close the connection with database at the end
	 */
	defer model.GetDB().Close()

	/*
		Run notification sender at background
	 */
	cnx, cancelNotifier := context.WithCancel(context.Background())
	go model.OnlyNotifyAboutGantaDeadline(cnx)
	defer cancelNotifier()

	/*
		creating a router instance
	 */
	var router = app.Create_new_invest_router()

	handlerOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"})

	/*
		port
	 */
	var port = "7000"
	var fname = "main"

	logr.WithFields(logr.Fields{
		"port": port,
	}).Info(fname)

	go http.ListenAndServe(":" + port, handlers.CORS(handlerOk, originOk, methodsOk)(router))

	/*
		ctrl + c -> shut down
	 */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("gracefully shutting down servers...")
}
