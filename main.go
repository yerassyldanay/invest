package main

import (
	"context"
	"fmt"
	logr "github.com/sirupsen/logrus"
	"invest/app"
	"invest/model"
	config "invest/utils/config"
	"invest/utils/helper"
	"invest/utils/logist"

	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var err error

	// environmental variables
	opts, err := config.LoadConfig("./environment/")
	helper.IfErrorPanic(err)
	_ = opts

	// Set up a connection with database
	_, err = model.EstablishDatabaseConnection(opts)
	helper.IfErrorPanic(err)

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

	// close the file at the end of the
	logist.InitiateLogFile()
	defer func() {
		fmt.Println("[FILE] closing connection...")
		err = logist.Get_file().Close()
		helper.IfErrorPanic(err)
	}()

	// this function stops a file rotator goroutine in the background
	defer func() {
		fmt.Println("[FILE-ROTATOR] closing connection...")
		logist.Get_file_rotator().Cancel <- true
	}()

	// run notification sender at background
	cnx, cancelNotifier := context.WithCancel(context.Background())
	go model.OnlyNotifyAboutGantaDeadline(cnx)
	defer cancelNotifier()

	// remove files at the background
	ctxRemoveAnalysisFiles, cancelRemoveAnalysisFilesCtx := context.WithCancel(context.Background())
	go model.Remove_files_left_after_analysis_periodically(ctxRemoveAnalysisFiles)
	defer cancelRemoveAnalysisFilesCtx()

	// creating a router instance
	var router = app.NewRouter()

	//handlerOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	//originOk := handlers.AllowedOrigins([]string{"*"})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"})

	/*
		port
	 */
	var port = "7000"
	var fname = "main"

	logr.WithFields(logr.Fields{
		"port": port,
	}).Info(fname)

	go http.ListenAndServe(":" + port, router)

	/*
		ctrl + c -> shut down
	 */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("gracefully shutting down servers...")
}
