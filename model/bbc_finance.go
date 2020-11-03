package model

import (
	"github.com/jinzhu/gorm"
	"invest/utils/constants"
)

type Finance struct {
	Id							uint64								`json:"id" gorm:"primary key"`
	ProjectId					uint64								`json:"project_id" gorm:"UNIQUE"`

	//TotalIncomeId				uint64								`json:"total_income_id"`
	TotalIncome					int						`json:"total_income"`

	//TotalProductionId			uint64								`json:"total_production_id"`
	TotalProduction				int						`json:"total_production"`

	//ProductionCostId			uint64								`json:"production_cost_id"`
	ProductionCost				int						`json:"production_cost"`

	//OperationalProfitId			uint64								`json:"operational_profit_id"`
	OperationalProfit			int						`json:"operational_profit"`

	//SettlementObligationsId			uint64					`json:"settlement_obligations_id"`
	SettlementObligations			int					`json:"settlement_obligations"`

	//OtherCostId					uint64								`json:"other_cost_id"`
	OtherCost					int						`json:"other_cost"`

	//PureProfitId				uint64							`json:"pure_profit_id"`
	PureProfit					int						`json:"pure_profit"`

	//TaxesId						uint64							`json:"taxes_id"`
	Taxes						int						`json:"taxes"`
}

//type FinanceCol struct {
//	Id						uint64					`json:"id" gorm:"primary key"`
//
//	Month					int64					`json:"month" gorm:"default:0"`
//	Year					int64					`json:"year" gorm:"default:0"`
//}

func (Finance) TableName() string {
	return "finances"
}

//func (FinanceCol) TableName() string {
//	return "finance_cols"
//}

/*
	update seq id before update & create operations
*/
func (fi *Finance) BeforeUpdate(tx *gorm.DB) error {
	return Update_sequence_id_thus_avoid_duplicate_primary_key_error(tx, constants.GormSeqIdFinance)
}

func(fi *Finance) BeforeCreate(tx *gorm.DB) error {
	return Update_sequence_id_thus_avoid_duplicate_primary_key_error(tx, constants.GormSeqIdFinance)
}

/*
	avoid the problem with having data not needed
*/
func (fic *Finance) AfterCreate(tx *gorm.DB) error {
	return nil
}

func (fi *Finance) Validate() error {
	if fi.ProjectId < 1 {
		return errorInvalidProjectId
	}

	return nil
}

func (fi *Finance) OnlyCreate(tx *gorm.DB) error {
	return tx.Create(fi).Error
}

func (fi *Finance) OnlySave(tx *gorm.DB) error {
	return tx.Save(fi).Error
}

func (fi *Finance) OnlyUpdateAll(tx *gorm.DB) error {
	return tx.Updates(*fi).Error
}

func (fi *Finance) OnlyGetByProjectId(tx *gorm.DB) (err error) {
	err = tx.First(fi, "project_id = ?", fi.ProjectId).Error
	return err
}

