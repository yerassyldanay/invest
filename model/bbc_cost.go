package model

type Cost struct {
	Id        uint64 `json:"id" gorm:"primary key"`
	ProjectId uint64 `json:"project_id" gorm:"UNIQUE"`

	//BuildingRepairId				uint64								`json:"building_and_repair_id"`
	BuildingRepairInvestor int `json:"building_repair_investor"`
	BuildingRepairInvolved int `json:"building_repair_involved"`

	//TechnologyEquipmentId			uint64								`json:"technology_equipment_id"`
	TechnologyEquipmentInvestor int `json:"technology_equipment_investor"`
	TechnologyEquipmentInvolved int `json:"technology_equipment_involved"`

	//WorkingCapitalId				uint64								`json:"working_capital_id"`
	WorkingCapitalInvestor int `json:"working_capital_investor"`
	WorkingCapitalInvolved int `json:"working_capital_involved"`

	//OtherCostId						uint64								`json:"other_cost_id"`
	OtherCostInvestor int `json:"other_cost_investor"`
	OtherCostInvolved int `json:"other_cost_involved"`

	//ShareInProjectId				uint64								`json:"share_in_project_id"`
	//ShareInProjectInvestor				int								`json:"share_in_project_investor"`
	//ShareInProjectInvolved				int								`json:"share_in_project_involved"`
}

//type CostCol struct {
//	Id						uint64							`json:"id" gorm:"primary key"`
//
//	InvestorFund				int								`json:"investor_fund"`
//	InvolvedFund				int								`json:"involved_fund"`
//}

func (Cost) TableName() string {
	return "costs"
}
