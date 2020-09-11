package model

import "time"

type AddInfo struct {
	Lang				string					`json:"lang" gorm:"-"`
	Created				time.Time				`json:"created" gorm:"default: now()"`
	Deleted				time.Time				`json:"deleted" gorm:"default: null"`
}

