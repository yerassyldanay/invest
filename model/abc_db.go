package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils/constants"
	"invest/utils/logist"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"os"
	"time"
)

//type DbConfigurations struct {
//
//}

var db *gorm.DB

/*
	E.g.
		postgresql://other@localhost/otherdb?connect_timeout=10&application_name=myapp
 */
func GetDbUri() (string, error) {

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

/*
	init it
 */
func Set_up_db() {

	var fname = "INIT_DB"
	var dbUri, err = GetDbUri()

	//fmt.Println(dbUri + "...")

	if err != nil {
		logist.SysMessage{
			FuncName: fname,
			Message:  err.Error(),
			Ok:       false,
			Lev:      constants.WarnLevel,
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
			logist.SysMessage{
				FuncName: fname,
				Message:  "connected to Postgres...",
				Ok:       true,
				Lev:      constants.InfoLevel,
			}.Log_system_message()
			break
		}

		logist.SysMessage{
			FuncName: fname,
			Message:  err.Error() + " sleeping for a while...",
			Ok:       false,
			Lev:      constants.WarnLevel,
		}.Log_system_message()
		time.Sleep(time.Second * constants.TimeSecToSleepBetweenDbConn)

		if i == constants.AttemptToConnectToDb {
			logist.SysMessage{
				FuncName: fname,
				Message:  "could not connect to db...",
				Ok:       false,
				Lev:      constants.WarnLevel,
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

	db.Debug().AutoMigrate(&Categor{}, &Comment{}, &Cost{}, &Document{},
		&Email{}, &Finance{}, &ForgetPassword{}, &Ganta{}, &Organization{},
		&Permission{}, &Phone{}, &Project{}, &Role{}, &SmtpServer{}, &SmtpHeaders{},
		&User{})

	db.Debug().AutoMigrate(&Notification{}, &NotificationInstance{}, &ProjectsUsers{})

	/*
		parameters of db
	 */
	db.DB().SetMaxOpenConns( constants.MaxNumberOpenConnToDb )

	err = PrepareSequenceId()
	if err != nil {
		fmt.Printf("could not prepare sequence id. err: ", err)
	}
}

func PrepareSequenceId() error {
	main_query := `
		select setval('costs_id_seq', (select max(id) from costs) + 1);
		select setval('finances_id_seq', (select max(id) from finances) + 1);
		select setval('gantas_id_seq', (select coalesce(max(id), 0) as id from gantas) + 1);
		select setval('emails_id_seq', (select max(id) from emails) + 1);
		select setval('phones_id_seq', (select max(id) from phones) + 1);
		select setval('users_id_seq', (select max(id) from users) + 1);
		select setval('roles_id_seq', (select max(id) from roles) + 1);
		select setval('projects_id_seq', (select max(id) from projects) + 1);
		select setval('organizations_id_seq', (select max(id) from organizations) + 1);
	`
	err := GetDB().Exec(main_query).Error
	return err
}

/*
	getter for gorm.DB object
 */
func GetDB () *gorm.DB {
	if db == nil {
		fmt.Printf("[CONN] Establishing a new db connection...")
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
