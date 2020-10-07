package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Analysis struct {
	Steps						[]int				`json:"steps"`
	Statuses					[]string				`json:"statuses"`
	Start						int64				`json:"start"`
	StartDate					time.Time			`json:"start_date"`
	End							int64				`json:"end"`
	EndDate						time.Time			`json:"end_date"`
}

func (a *Analysis) Get_projects_by_steps(offset interface{}, tx *gorm.DB) (projects []Project, err error) {
	err = tx.Find(&projects, "step in (?) and status in (?) and created > ? and created < ?",
		a.Steps, a.Statuses, a.StartDate, a.EndDate).Offset(offset).Limit(GetLimit + 10).Error
	return projects, err
}
