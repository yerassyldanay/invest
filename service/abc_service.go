package service

type InvestService struct {
	BasicInfo
}

type BasicInfo struct {
	UserId				uint64				`json:"user_id"`
	RoleId				uint64				`json:"role_id"`
	Lang				string				`json:"lang"`
}

