package model

import (
	"github.com/joho/godotenv"
	"invest/utils/helper"
	"os"
	"path/filepath"
	"regexp"
)

func Load_env_values() error {
	current_path, _ := os.Getwd()

	var env1_general, env2_sendgrid string
	if ok, err := regexp.Match("[a-zA-Z0-9]*/invest/+", []byte(current_path)); ok && err == nil {
		env1_general = "../env/.env"
		env2_sendgrid = "../env/sendgrid.env"
	} else {
		//fmt.Println("err: ", err)
		env1_general = "./env/.env"
		env2_sendgrid = "./env/sendgrid.env"
	}

	/*
		the following call loads all env variables in the .env file
	*/
	path_env, err1 := filepath.Abs(env1_general)
	path_sendgrid, err2 := filepath.Abs(env2_sendgrid)

	if err1 != nil || err2 != nil {
		err := helper.If_condition_then(err1 != nil, err1, err2).(error)
		return err
	}

	//fmt.Println("path_* : ", path_env, path_sendgrid, env1_general, env2_sendgrid)

	if err := godotenv.Load(path_env, path_sendgrid); err != nil {
		//fmt.Println(err.Error())
		return err
	}

	return nil
}