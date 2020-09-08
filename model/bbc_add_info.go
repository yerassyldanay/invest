package model

import "time"

type AddInfo struct {
	Lang				string					`json:"lang"`
	Created				time.Time				`json:"created" gorm:"default: now()"`
	Deleted				time.Time				`json:"deleted" gorm:"default: null"`
}

