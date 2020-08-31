package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

type Finance struct {
	Id						uint64								`json:"id" gorm:"primary key"`

	ProjectId				uint64								`json:"project_id" gorm:"UNIQUE"`
	Project					Project								`json:"project" gorm:"foreignkey:ProjectId"`
	
	LandId					uint64									`json:"land_id"`
	Land					FinanceCol								`json:"land" gorm:"foreignkey:LandId"`

	TechId					uint64									`json:"tech_id"`
	Tech					FinanceCol								`json:"tech" gorm:"foreignkey:TechId"`
	
	CapitalId				uint64									`json:"capital_id"`
	Capital					FinanceCol								`json:"capital" gorm:"foreignkey:CapitalId"`
	
	OtherId					uint64									`json:"other_id"`
	Other					FinanceCol								`json:"other" gorm:"foreignkey:OtherId"`
	
	SumId					uint64									`json:"sum_id"`
	Sum						FinanceCol								`json:"sum" gorm:"foreignkey:SumId"`
}

type FinanceCol struct {
	Id						uint64									`json:"id" gorm:"primary key"`

	Sum						int64									`json:"sum" gorm:"default:0"`
	SPK						int64									`json:"spk" gorm:"default:0"`
	Initiator				int64									`json:"initiator" gorm:"default:0"`
	Funded					int64									`json:"funded" gorm:"default:0"`
}

func (Finance) TableName() string {
	return "finances"
}

func (FinanceCol) TableName() string {
	return "finance_cols"
}

/*
	update seq id before update & create operations
 */
func (fi *Finance) BeforeUpdate(tx *gorm.DB) error {
	return Update_sequence_id_thus_avoid_duplicate_primary_key_error(tx, utils.GormSeqIdFinance)
}

func(fi *Finance) BeforeCreate(tx *gorm.DB) error {
	return Update_sequence_id_thus_avoid_duplicate_primary_key_error(tx, utils.GormSeqIdFinance)
}

/*
	avoid the problem with having data not needed
 */
func (fic *Finance) AfterCreate(tx *gorm.DB) error {
	var main_query = `delete from finance_cols where id not in (
		select land_id from finances
			union
		select capital_id from finances
			union
		select other_id from finances
			union
		select sum_id from finances
			union
		select tech_id from finances
	);`

	return tx.Exec(main_query).Error
}
