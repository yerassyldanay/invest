package model

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"regexp"
)

func Load_env_values() error {
	current_path, _ := os.Getwd()

	var env1_general string
	if ok, err := regexp.Match("[a-zA-Z0-9]*/invest/+", []byte(current_path)); ok && err == nil {
		env1_general = "../env/.env"
	} else {
		env1_general = "./env/.env"
	}

	/*
		the following call loads all env variables in the .env file
	*/
	path_env, err := filepath.Abs(env1_general)

	if err != nil {
		return err
	}

	if err := godotenv.Load(path_env); err != nil {
		return err
	}

	return nil
}