package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
	"time"
)

type SmtpServer struct {
	Id					uint64					`json:"id" gorm:"primaryKey"`
	Host				string					`json:"host"`
	Port				int						`json:"port"`
	Username			string					`json:"username"`
	Password			string					`json:"password"`
	LastUsed			time.Time				`json:"last_used" gorm:"now()"`
}

func (ss *SmtpServer) TableName() string {
	return "smtp_servers"
}

func (ss *SmtpServer) OnlyCreate(tx *gorm.DB) error {
	err := tx.Create(ss).Error
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
	err := tx.Model(&SmtpServer{Id: ss.Id}).Update("last_used", utils.GetCurrentTime()).Error
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
