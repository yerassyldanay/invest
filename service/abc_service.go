package service

import "time"

type InvestService struct {
	TimeUTC				time.Time			`json:"time_utc"`
	Time				time.Time			`json:"time" `
	Offset				interface{}
	BasicInfo
}

type BasicInfo struct {
	UserId				uint64				`json:"user_id"`
	RoleId				uint64				`json:"role_id"`
	RoleName			string				`json:"role_name"`
	Lang				string				`json:"lang"`
}

