package model

type ProjectStatsRaw struct {
	Status				string				`json:"status"`
	Number				uint				`json:"number"`
}

type ProjectStatsOnStatuses struct {
	Newone				uint				`json:"newone"`
	Inprogress			uint				`json:"inprogress"`
	Done				uint				`json:"done"`
	Rejected			uint				`json:"rejected"`
}
