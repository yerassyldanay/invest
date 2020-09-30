package utils

import (
	"fmt"
	logr "github.com/sirupsen/logrus"

	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

/*
	this helps rotating log file while writing into them
	problem is that when files gets large, it becomes hard to open and read (take long time)
 */
type FileRotator struct {
	LastDate			time.Time
	FileName			string
	Cancel				chan bool
	Wait				sync.WaitGroup
	File				*os.File
}

/*
	instance of it
		chan is used to notify about time to change the name of the file
 */
var FR = FileRotator{
	LastDate: 			time.Now(),
	FileName: 			"",
	Cancel:    			make(chan bool, 1),
	Wait:				sync.WaitGroup{},
	File: 				nil,
}

/*
	create a file object, which allows to close at the end
	to escape resource leak
 */
var file *os.File

/*
	file name constructor
 */
func (fr *FileRotator) Construct_file_name() string {
	y, m, d := time.Now().Date()
	var temp = fmt.Sprintf("log_%d_%d_%d.log", y, m, d)
	return temp
}

/*
	this function will be running in the background
	and it rotates files based on the date
	for each day => one log file
 */
func Rotate_log_file_periodically(fr *FileRotator) {

	var fname = "fn_rotate_log_files"

	defer fr.Wait.Done()
	select {
	case <- time.Tick(time.Hour * 12):
		temp := fr.Construct_file_name()
		fr.FileName = temp
		fr.Set_log_file()
	case <- fr.Cancel:
		SysMessage {
			FuncName: fname,
			Message:  "finishing the rotate fr goroutine",
			Lev:      InfoLevel,
			Ok:       true,
		}.Log_system_message()
		return
	}
}

/*
	create a file if there is not any which matches the file name
	which is associated with today
 */
func (fr *FileRotator) Set_log_file() {

	if fr.FileName == "" {
		temp := fr.Construct_file_name()
		fr.FileName = temp
	}

	var err error
	var file_path string

	/*
		this is to handle the mistake arising when running this function from different paths
			main.go - /invest
			intest - /invest/intest
	 */
	current_path, _ := os.Getwd()
	if strings.Contains(current_path, "/invest/intest") {
		file_path, err = filepath.Abs("../" + FolderLogFiles + "/")
		if err != nil {
			log.Println(err)
		}
	} else {
		file_path, err = filepath.Abs("./" + FolderLogFiles + "/")
		if err != nil {
			log.Println(err)
		}
	}

	file_path = filepath.Join(file_path, FR.FileName)
	//fmt.Println("file_path: ", file_path, current_path)

	file, err = os.OpenFile(file_path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	fr.File = file

	/*
		It is not good idea to close the file within the scope of this function
		This issue is handled using GetLogFile() (at the end of the file) and closing the file within main function
	*/
	//defer file.Close()

	/*
	The format would be as following:
		{"level":"fatal","msg":"The ice breaks!","number":100,"omg":true, "time":"2014-03-10 19:57:38.562543128 -0400 EDT"}
	 */
	logr.SetFormatter(&logr.TextFormatter{
		TimestampFormat: 	time.RFC3339,
		DisableColors: 		true,
	})

	/*
		set the output reader
	 */
	mw := io.MultiWriter(file, os.Stdout)
	logr.SetOutput(mw)

	/*
		you can set the level, which (and above levels) will be written to the file
	 */
	//logr.SetLevel(logr.FatalLevel)
}

/*
	initiate the log file
 */
func init() {
	fmt.Println("logging ...")
	Get_file_rotator().Set_log_file()

	/*
		this will check & rotate the log file periodically
	*/
	Get_file_rotator().Wait.Add(1)
	go Rotate_log_file_periodically(Get_file_rotator())
}

/*
	getter
 */
func Get_file_rotator () (*FileRotator) {
	return &FR
}

/*
	getter for file object
 */
func Get_file() *os.File {
	var fname = "GET_FILE"
	SysMessage {
		FuncName: fname,
		Message:  "closing the file...",
		Lev:      InfoLevel,
		Ok:       true,
	}.Log_system_message()
	return file
}

