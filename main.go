package main

import (
	"fmt"
	logr "github.com/sirupsen/logrus"
	"invest/app"
	"invest/model"
	"invest/utils"
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

	//var b map[string][]string
	//_ = json.Unmarshal([]byte(`{"eng":["aa","bb"]}`), &b)
	//fmt.Println(b)
	//fmt.Println(b["eng"])

	/*
		migration
	*/
	if err := model.Migration(); err != nil {
		utils.SysMessage{
			FuncName: "MAIN",
			Message:  "migration: " + err.Error(),
			Ok:       false,
			Lev:      utils.WarnLevel,
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
	defer utils.Get_file()

	/*
		this function stops the goroutines in the background
	 */
	defer func() {
		utils.Get_file_rotator().Cancel <- true
	}()

	/*
		close the connection with database at the end
	 */
	defer model.GetDB().Close()

	/*
		prepare permissions table
	 */
	app.Prepare_permissions()

	/*
		creating a router instance
	 */
	var router = app.Create_new_invest_router()

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
