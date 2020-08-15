package model

import "time"

/*
	this lines of code are repeated again and again
 */
type Model struct {
	Id					uint64				`json:"-" gorm:"AUTO_INCREMENT; primary_key"`
	DeletedAt			time.Time			`json:"-" gorm:"default: null"`
}
