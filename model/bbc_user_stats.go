package model

type UserStatsRaw struct {
	Status				string 				`json:"name"`
	Number				uint				`json:"count"`
}

type UserStats struct {
	Newone				uint				`json:"newone"`
	Inprogress			uint				`json:"inprogress"`
	Rejected			uint				`json:"rejected"`
	Done				uint				`json:"done"`
}

type UserStatsRawList []UserStatsRaw
