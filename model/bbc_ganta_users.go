package model

type GantaStepResp struct {
	StepId				uint64						`json:"step_id" gorm:""`
	Statuses			[]GantaStepRespStatus
}

type GantaStepRespStatus struct {
	StepId						uint64				`json:"step_id"`
	UserId						uint64				`json:"user_id"`
	Status						string				`json:"status" default:"newone"`
}

