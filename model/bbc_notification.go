package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

/*
	this is the struct of notifications message which will be stored on database

	NotificationInstance will carry:
		* receiver address
		* notification body (Notification), which will be stored on another table
 */
type NotificationInstance struct {
	ToAddress			string				`json:"to_address"`
	NotificationId		uint64				`json:"notification_id"`
	Notification		Notification		`json:"notification" gorm:"foreignKey:notifications.id"`
}

/*
	Notification will bear only body of the message
		* it will be stored only once
		* its id will be used in NotificationInstance table
 */
type Notification struct {
	Id						uint64				`json:"id" gorm:"primaryKey"`
	// who send the message
	FromAddress				string				`json:"from_address"`
	// the notification is connected / related to which project
	ProjectId				uint64				`json:"project_id" gorm:"default:0"`
	// body
	Html					string				`json:"html"`
	// body in text/plain
	Plain					string				`json:"plain"`
	// created date
	Created					time.Time			`json:"created" gorm:"default:now()"`
}

// NotificationInstance
func (ni *NotificationInstance) TableName() string {
	return "notification_instances"
}

// Notification
func (n *Notification) TableName() string {
	return "notifications"
}

// create notification instance
func (ni *NotificationInstance) OnlyCreate(tx *gorm.DB) (error) {
	err := tx.Create(ni).Error
	return err
}

// create n
func (n *Notification) OnlyCreate(tx *gorm.DB) error {
	err := tx.Create(n).Error
	return err
}

// get n-s
func (ni *NotificationInstance) OnlyGetNotificationsByEmailAndProjectId(address string, project_id uint64, offset interface{}, tx *gorm.DB) ([]Notification, error) {
	notifications := []Notification{}
	err := tx.Raw("select n.* from notification_instances ni join notifications n on ni.notification_id = n.id " +
		" where ni.to_address = ? and n.project_id = ? order by n.created desc limit 30 offset ?; ", address, project_id, offset).
		Scan(&notifications).Error
	return notifications, err
}
