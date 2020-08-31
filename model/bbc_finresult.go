package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils"
)

type Finresult struct {
	Id							uint64								`json:"id" gorm:"primary key"`

	ProjectId					uint64								`json:"project_id" gorm:"UNIQUE"`
	Project						Project								`json:"project" gorm:"foreignkey:ProjectId"`

	TotalIncomeId				uint64								`json:"total_income_id"`
	TotalIncome					FinresultCol						`json:"total_income"`

	TotalProductionId			uint64								`json:"total_production_id"`
	TotalProduction				FinresultCol						`json:"total_production"`

	ProductionCostId			uint64								`json:"production_cost_id"`
	ProductionCost				FinresultCol						`json:"production_cost"`

	OperationalProfitId			uint64								`json:"operational_profit_id"`
	OperationalProfit			FinresultCol						`json:"operational_profit"`

	CancellationId				uint64							`json:"cancellation_id"`
	Cancellation				FinresultCol					`json:"cancellation"`

	OtherCostId					uint64							`json:"other_cost_id"`
	OtherCost					FinresultCol							`json:"other_cost"`

	PrivateCostId				uint64							`json:"private_cost_id"`
	PrivateCost					FinresultCol					`json:"private_cost"`
}

type FinresultCol struct {
	Id						uint64					`json:"id" gorm:"primary key"`

	Month					int64					`json:"month" gorm:"default:0"`
	Year					int64					`json:"year" gorm:"default:0"`
	Total					int64					`json:"total" gorm:"default:0"`
}

func (Finresult) TableName() string {
	return "finresults"
}

func (FinresultCol) TableName() string {
	return "finresult_cols"
}

/*
	update seq id before update & create operations
*/
func (fi *Finresult) BeforeUpdate(tx *gorm.DB) error {
	return Update_sequence_id_thus_avoid_duplicate_primary_key_error(tx, utils.GormSeqIdFinance)
}

func(fi *Finresult) BeforeCreate(tx *gorm.DB) error {
	return Update_sequence_id_thus_avoid_duplicate_primary_key_error(tx, utils.GormSeqIdFinance)
}

/*
	avoid the problem with having data not needed
*/
func (fic *Finresult) AfterCreate(tx *gorm.DB) error {
	var main_query = `delete from finresult_cols where id not in (
		select total_income_id from finresults
		union
		select total_production_id from finresults
		union
		select production_cost_id from finresults
		union
		select operational_profit_id from finresults
		union
		select cancellation_id from finresults
		union
		select other_cost_id from finresults
		union
		select private_cost_id from finresults
	);`

	return tx.Exec(main_query).Error
}


