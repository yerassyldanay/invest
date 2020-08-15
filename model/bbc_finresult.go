package model

type Finresult struct {
	Id							uint64								`json:"id" gorm:"primary key"`

	ProjectId					uint64								`json:"project_id"`
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