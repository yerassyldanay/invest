package model

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	//"github.com/pressly/goose"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"invest/utils"
	"os"
	"path/filepath"
	"time"
)

//type DbConfigurations struct {
//
//}

var db *gorm.DB

/*
	this allows to choose the database and set parameters
 */
func chooseDbDriver(dbtype, dbpath string) goose.DBDriver {
	drive := goose.DBDriver{
		Name:    		dbtype,
		OpenStr: 		dbpath,
		Import:  		"",
		Dialect: 		nil,
	}

	switch dbtype {
	default:
		drive.Import = "github.com/lib/pq"
		drive.Dialect = &goose.PostgresDialect{}
	}

	return drive
}

/*
	E.g.
		postgresql://other@localhost/otherdb?connect_timeout=10&application_name=myapp
 */
func Get_db_uri() (string, error) {

	if err := Load_env_values(); err != nil {
		return "", err
	}

	/*
		the env variables are loaded above
	*/
	var dbUsername = os.Getenv("POSTGRES_USER")
	var dbPassword = os.Getenv("POSTGRES_PASSWORD")
	var dbName = os.Getenv("POSTGRES_DB")
	var dbHost = os.Getenv("POSTGRES_HOST")
	var dbPort = os.Getenv("POSTGRES_PORT")

	var dbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUsername, dbName, dbPassword)

	//fmt.Println("dbUri: ", dbUri)

	return dbUri, nil
}

func Migration() error {
	//var dbName = os.Getenv("POSTGRES_DB")
	var dbUri, err = Get_db_uri()

	if err != nil {
		return err
	}

	migrateConf := &goose.DBConf {
		MigrationsDir: 		os.Getenv("MIGRATION_PATH"),
		Env:           		os.Getenv("ENV"),
		Driver:        		chooseDbDriver("postgres", dbUri),
	}

	abs, _ := filepath.Abs(migrateConf.MigrationsDir)
	//fmt.Println(abs)

	latest, err := goose.GetMostRecentDBVersion(abs)
	if err != nil {
		return err
	}

	//fmt.Println(latest)
	var tdb = GetDB().DB()
	if tdb == nil {
		return errors.New("*DB is nil")
	}

	err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, tdb)
	if err != nil {
		return err
	}

	return nil
}

/*
	init it
 */
func Set_up_db() {

	var fname = "INIT_DB"
	var dbUri, err = Get_db_uri()

	//fmt.Println(dbUri + "...")

	if err != nil {
		utils.SysMessage{
			FuncName: fname,
			Message:  err.Error(),
			Ok:       false,
			Lev:      utils.WarnLevel,
		}.Log_system_message()
		return
	}

	/*
		actual connection to the database
	 */
	var i = 0

	for {
		db, err = gorm.Open("postgres", dbUri)
		if err == nil {
			utils.SysMessage{
				FuncName: fname,
				Message:  "connected to Postgres...",
				Ok:       true,
				Lev:      utils.InfoLevel,
			}.Log_system_message()
			break
		}

		utils.SysMessage{
			FuncName: fname,
			Message:  err.Error() + " sleeping for a while...",
			Ok:       false,
			Lev:      utils.WarnLevel,
		}.Log_system_message()
		time.Sleep(time.Second * utils.TimeSecToSleepBetweenDbConn)

		if i == utils.AttemptToConnectToDb {
			utils.SysMessage{
				FuncName: fname,
				Message:  "could not connect to db...",
				Ok:       false,
				Lev:      utils.WarnLevel,
			}.Log_system_message()
			return
		}

		i++
	}

	/*
		the following call makes changes to the database based on the changes in provided struct-s
	 */
	//db.Debug().AutoMigrate(&Admin{}, &Category{}, &Company{}, &CivilServant{}, &Email{},
	//&Investor{}, InvestorAndCompany{}, &Phone{}, &Position{}, &Project{}, &ProjectDoc{},
	//&ProjectCivilConnection{}, &SendgridMessage{})

	db.Debug().AutoMigrate(&Categor{}, &Comment{}, &Document{}, &Email{}, &Finance{}, &FinanceCol{},
		&Finresult{}, FinresultCol{}, &Ganta{}, &Organization{}, &Permission{},
		&Phone{}, &Project{}, &ProjectStatus{}, &Role{}, &SendgridMessage{}, &SendgridMessageStore{},
		&User{})

	db.Debug().AutoMigrate(&ProjectsUsers{})

	/*
		parameters of db
	 */
	db.DB().SetMaxOpenConns(utils.MaxNumberOpenConnToDb)

	/*
		these lines of code will make sure that everytime we create a row
			sequence id will be updated
		as there is a problem with sync of seq. id.
	 */
	db.Callback().Create().Before("gorm:update").Register("updateSeqId", func(scope *gorm.Scope) {
		fmt.Println("updating hook is working: ", scope.TableName())
		_ = Update_sequence_id_thus_avoid_duplicate_primary_key_error(scope.DB(), "default")
	})
}

/*
	getter for gorm.DB object
 */
func GetDB () *gorm.DB {
	if db == nil {
		Set_up_db()
	}
	return db
}

/*
	rollback at the end if something happens
 */
func Rollback(trans *gorm.DB) {
	if trans != nil {
		trans.Rollback()
	}
}