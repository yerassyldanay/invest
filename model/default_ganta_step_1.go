package model

import (
	"invest/utils"
)

/*
	Ganta possesses a field called 'isHidden':
 */
var DefaultGantaParentsOfStep1 = []Ganta{
	{
		IsAdditional:   false,
		Kaz:            "Құжаттарды жүктеу",
		Rus:            "Загрузка документов инвестором",
		Eng:            "Uploading files by an investor",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,

		Responsible:    utils.RoleInvestor,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Менеджерді тағайындау",
		Rus:            "Назначение менеджера",
		Eng:            "Assignment of a manager",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingAdmin,

		Responsible:    utils.RoleAdmin,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Құжаттарды тексеру",
		Rus:            "Проверка документов",
		Eng:            "Review of documents",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingManager,

		Responsible:    utils.RoleManager,
		IsDone: 		false, // always false
		IsDocCheck: 	true,
	},
	{
		IsAdditional:   false,
		Kaz:            "Инвестициялық комитет басшысының тексеруі",
		Rus:            "Проверка главой инвестиционного комитета",
		Eng:            "Review by a head of an investment committee",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingAdmin,

		Responsible:    utils.RoleAdmin,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Қаржы-қ / құқық-қ сараптама",
		Rus:            "Фин / юр эспертиза",
		Eng:            "A finance / jurisdiction expertise",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingExpert,
		Responsible:    utils.RoleExpert,
		IsDocCheck: 	true,
	},
	{
		IsAdditional:   false,
		Kaz:            "Инвестициялық комитет басшысының тексеруі",
		Rus:            "Проверка главой инвестиционного комитета",
		Eng:            "Review by a head of an investment committee",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingAdmin,

		Responsible:    utils.RoleAdmin,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "ИК-тің қарауында",
		Rus:            "Рассмотрение на ИК СПК",
		Eng:            "Review by an invest committee",
		DurationInDays: 6,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvCommittee,
		Responsible:    utils.RoleManager,
	},
	{
		IsAdditional:   false,
		Kaz:            "ХҚО-на сұраным жолдау",
		Rus:            "Заявки в ЦОН",
		Eng:            "Application to People Service Center",
		DurationInDays: 30 * 6,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusRegistrationOfLandPlot,
		Responsible:    utils.RoleManager,
	},
	{
		IsAdditional:   false,
		Kaz:            "Смета (жоба) құжаттары",
		Rus:            "Предоставление в СПК ПСД с экспертизой, ТЭО",
		Eng:            "Provision of a design & estimate doc.",
		DurationInDays: 9 * 30,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingManager,
		Responsible:    utils.RoleManager,
	},
}
