package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/yerassyldanay/invest/utils/config"
	"github.com/yerassyldanay/invest/utils/constants"
)

var db *gorm.DB

// EstablishDatabaseConnection establishes connection with database
func EstablishDatabaseConnection(opts config.Config) (*gorm.DB, error) {
	// connect
	tempDb, err := gorm.Open("postgres", opts.GetDatabaseSource())
	if err != nil {
		return tempDb, fmt.Errorf("failed to establish connection with database. err: %v", err)
	}
	db = tempDb

	// 		parameters of database
	db.DB().SetMaxOpenConns(constants.MaxNumberOpenConnToDb)

	//db.LogMode(true)

	return tempDb, err
}

// GetDB getter for gorm.DB object
func GetDB() *gorm.DB {
	if db == nil {
		fmt.Printf("[CONN] Establishing a new database connection...")
		opts, err := config.LoadConfig("../environment/.local.env")
		if err != nil {
			fmt.Printf("failed to get database connection. err: %v\n", err)
		}
		db, err = EstablishDatabaseConnection(opts)
		if err != nil {
			fmt.Printf("failed to get database connection. err: %v\n", err)
		}
	}
	return db
}

// Rollback rollback at the end if something happens
func Rollback(trans *gorm.DB) {
	if trans != nil {
		trans.Rollback()
	}
}

func HelperPrint(any interface{}) {
	b, err := json.MarshalIndent(any, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
