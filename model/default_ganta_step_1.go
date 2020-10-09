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
		Kaz:            "Загрузка документов Инициатором проекта",
		Rus:            "Загрузка документов Инициатором проекта",
		Eng:            "Загрузка документов Инициатором проекта",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,

		Responsible:    utils.RoleInvestor,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Прикрепление сотрудников СПК к проекту",
		Rus:            "Прикрепление сотрудников СПК к проекту",
		Eng:            "Прикрепление сотрудников СПК к проекту",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingAdmin,

		Responsible:    utils.RoleAdmin,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Рассмотрение заявки / документов",
		Rus:            "Рассмотрение заявки / документов",
		Eng:            "Рассмотрение заявки / документов",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingManager,

		Responsible:    utils.RoleManager,
		IsDone: 		false, // always false
		IsDocCheck: 	true,
	},
	{
		IsAdditional:   false,
		Kaz:            "Рассмотрение главой инвестиционного департамента",
		Rus:            "Рассмотрение главой инвестиционного департамента",
		Eng:            "Рассмотрение главой инвестиционного департамента",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingAdmin,

		Responsible:    utils.RoleAdmin,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Предварительная экспертиза (ДИП, ДЭП, ЮД, ДБ и НУ)",
		Rus:            "Предварительная экспертиза (ДИП, ДЭП, ЮД, ДБ и НУ)",
		Eng:            "Предварительная экспертиза (ДИП, ДЭП, ЮД, ДБ и НУ)",
		DurationInDays: 7,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingExpert,
		Responsible:    utils.RoleExpert,
		IsDocCheck: 	true,
	},
	{
		IsAdditional:   false,
		Kaz:            "Рассмотрение Заявки на ИК СПК",
		Rus:            "Рассмотрение Заявки на ИК СПК",
		Eng:            "Рассмотрение Заявки на ИК СПК",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvCommittee,
		Responsible:    utils.RoleManager,
	},
	{
		IsAdditional:   false,
		Kaz:            "Протокол Решения ИК",
		Rus:            "Протокол Решения ИК",
		Eng:            "Протокол Решения ИК",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingBoard,
		Responsible:    utils.RoleManager,
	},
	{
		IsAdditional:   false,
		Kaz:            "Заявки в ЦОН",
		Rus:            "Заявки в ЦОН",
		Eng:            "Заявки в ЦОН",
		DurationInDays: 5,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusRegistrationOfLandPlot,
		Responsible:    utils.RoleManager,
	},
	{
		IsAdditional:   false,
		Kaz:            "Предоставление в СПК ПСД с экспертизой, ТЭО*",
		Rus:            "Предоставление в СПК ПСД с экспертизой, ТЭО*",
		Eng:            "Предоставление в СПК ПСД с экспертизой, ТЭО*",
		DurationInDays: 5,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingManager,
		Responsible:    utils.RoleManager,
	},
}
