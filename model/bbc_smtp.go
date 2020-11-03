package model

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"invest/utils/helper"
	"time"
)

type SmtpServer struct {
	Id					uint64					`json:"id" gorm:"primaryKey"`
	Host				string					`json:"host"`
	Port				int						`json:"port"`
	Username			string					`json:"username"`
	Password			string					`json:"password"`
	LastUsed			time.Time				`json:"last_used" gorm:"now()"`
	Headers				[]SmtpHeaders			`json:"headers"`
}

// these headers are not used at this moment,
// but might be needed later
type SmtpHeaders struct {
	Id							uint64					`json:"id" gorm:"primaryKey"`
	SmtpServerId				uint64					`json:"smtp_server_id"`
	Key							string					`json:"header"`
	Value						string					`json:"value"`
}

func (ss *SmtpServer) TableName() string {
	return "smtp_servers"
}

// errors
var errorSmtpInvalidPortNumber = errors.New("invalid port number")
var errorSmtpInvalidHost = errors.New("invalid host")
var errorSmtpInvalidUsernameOrPassword = errors.New("invalid username or password")

// validate smtp server
func (ss *SmtpServer) Validate() error {
	switch {
	case ss.Port < 1:
		return errorSmtpInvalidPortNumber
	case len(ss.Host) == 0:
		return errorSmtpInvalidHost
	case len(ss.Username) < 1 || len(ss.Password) < 1:
		return errorSmtpInvalidUsernameOrPassword
	}

	return nil
}

func (ss *SmtpServer) OnlyCreate(tx *gorm.DB) error {
	err := tx.Create(ss).Error
	return err
}

func (ss *SmtpServer) OnlyUpdateFieldsById(tx *gorm.DB, fields ... interface{}) (error) {
	err := tx.Where("id = ?", ss.Id).Select(fields).Updates(ss).Error
	return err
}

func (ss *SmtpServer) OnlyUpdateUsernameAndPasswordById(tx *gorm.DB) error {
	err := tx.Model(&SmtpServer{Id: ss.Id}).Updates(map[string]interface{}{
		"username": ss.Username,
		"password": ss.Password,
	}).Error
	return err
}

// update last used time
func (ss *SmtpServer) OnlyUpdateLastTimeUsed(tx *gorm.DB) (error) {
	err := tx.Model(&SmtpServer{Id: ss.Id}).Update("last_used", helper.GetCurrentTime()).Error
	return err
}

// only delete
func (ss *SmtpServer) OnlyDeleteById(tx *gorm.DB) (error) {
	err := tx.Delete(ss, "id = ?", ss.Id).Error
	return err
}

// get by id
func (ss *SmtpServer) OnlyGetById(tx *gorm.DB) (error) {
	err := tx.First(ss, "id = ?", ss.Id).Error
	return err
}

// get all
func (ss *SmtpServer) OnlyGetAll(tx *gorm.DB) ([]SmtpServer, error) {
	smtps := []SmtpServer{}
	err := tx.Preload("Headers").Find(&smtps).Error
	return smtps, err
}

// get by host & port
func (ss *SmtpServer) OnlyGetByHostAndPort(tx *gorm.DB) (error) {
	err := tx.First(ss, "host = ? and port = ?", ss.Host, ss.Port).Error
	return err
}

// only get one
func (ss *SmtpServer) OnlyGetOne(tx *gorm.DB) (error) {
	err := tx.Raw("select * from smtp_servers order by last_used desc;").Scan(ss).Error
	return err
}

func (ss *SmtpServer) OnlySaveById(tx *gorm.DB) (error) {
	err := tx.Where("id = ?", ss.Id).Save(ss).Error
	return err
}

// only set headers
func (ss *SmtpServer) OnlySetHeaders(tx *gorm.DB) (error) {
	if len(ss.Headers) < 1 {
		return nil
	}

	main_query := bytes.Buffer{}
	main_query.WriteString("insert into smtp_headers (smtp_server_id, key, value) values")
	for i, header := range ss.Headers {
		if i != 0 {
			main_query.WriteString(" , ")
		}

		// check key
		if err := helper.OnlyCheckSqlInjection(header.Key); err != nil {
			return err
		}

		// check value
		if err := helper.OnlyCheckSqlInjection(header.Value); err != nil {
			return err
		}

		main_query.WriteString(fmt.Sprintf(" (%d, '%s', '%s') ", ss.Id, header.Key, header.Value))
	}

	main_query.WriteString(" ; ")
	a := main_query.String()
	_ = a

	if n := tx.Exec(main_query.String()).RowsAffected; n <= 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (ss *SmtpServer) OnlyDeleteAllHeadersBySmtpId(tx *gorm.DB) (error) {
	err := tx.Delete(&SmtpHeaders{}, "smtp_server_id = ?", ss.Id).Error
	_ = err
	return nil
}
