package model

import (
	"invest/utils"

	"time"
)

var DefaultGantaStep1Time = utils.GetCurrentTime()
var Day = time.Hour * 24

/*
	Ganta possesses a field called 'isHidden':
 */
var DefaultGantaParentsOfStep1 = []Ganta{
	{
		IsAdditional:   false,
		Kaz:            "Загрузка документов Инициатором проекта",
		Rus:            "Загрузка документов Инициатором проекта",
		Eng:            "Загрузка документов Инициатором проекта",
		DurationInDays: 4,
		GantaChildren:  DefaultGantaChildren1,

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
		DurationInDays: 1,
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
		DurationInDays: 1,
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
		DurationInDays: 1,
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
		DurationInDays: 6,
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
		DurationInDays: 5,
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
		DurationInDays: 5,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingBoard,
		Responsible:    utils.RoleManager,
	},
	{
		IsAdditional:   false,
		Kaz:            "Оформление ДИП документов на право землепользования",
		Rus:            "Оформление ДИП документов на право землепользования",
		Eng:            "Оформление ДИП документов на право землепользования",
		DurationInDays: 5,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusRegistrationOfLandPlot,
		Responsible:    utils.RoleManager,
	},
	{
		IsAdditional:   false,
		Kaz:            "Предоставление в СПК ПСД с экспертизой",
		Rus:            "Предоставление в СПК ПСД с экспертизой",
		Eng:            "Предоставление в СПК ПСД с экспертизой",
		DurationInDays: 5,
		GantaChildren:  []Ganta{},

		Step:           1,
		Status: 		utils.ProjectStatusPendingManager,
		Responsible:    utils.RoleManager,
	},
}

var DefaultGantaChildren1 = []Ganta{
	{
		IsAdditional:   false,
		Kaz:            "Свидетельство/справка о государственной регистрации юридического лица",
		Rus:            "Свидетельство/справка о государственной регистрации юридического лица",
		Eng:            "Свидетельство/справка о государственной регистрации юридического лица",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,                                 // will be set

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,  // this status will not be considered instead consider the status of the child
	},
	{
		IsAdditional:   false,
		Kaz:            "Устав, Учредительный договор",
		Rus:            "Устав, Учредительный договор",
		Eng:            "Устав, Учредительный договор",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Решение Уполномоченного органа юридического лица о назначении исполнительного органа",
		Rus:            "Решение Уполномоченного органа юридического лица о назначении исполнительного органа",
		Eng:            "Решение Уполномоченного органа юридического лица о назначении исполнительного органа",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Копии документов, удостоверяющих личность руководителя, главного бухгалтера предприятия, всех учредителей предприятия, владеющих более 10% уставного капитала;",
		Rus:            "Копии документов, удостоверяющих личность руководителя, главного бухгалтера предприятия, всех учредителей предприятия, владеющих более 10% уставного капитала;",
		Eng:            "Копии документов, удостоверяющих личность руководителя, главного бухгалтера предприятия, всех учредителей предприятия, владеющих более 10% уставного капитала;",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Документы, подтверждающие наличие денежных средств для реализации проекта",
		Rus:            "Документы, подтверждающие наличие денежных средств для реализации проекта",
		Eng:            "Документы, подтверждающие наличие денежных средств для реализации проекта",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Отчет Первого кредитного бюро",
		Rus:            "Отчет Первого кредитного бюро",
		Eng:            "Отчет Первого кредитного бюро",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Справка из государственного органа об отсутствии задолженности по налогам и обязательным платежам в бюджет, либо акт сверки",
		Rus:            "Справка из государственного органа об отсутствии задолженности по налогам и обязательным платежам в бюджет, либо акт сверки",
		Eng:            "Справка из государственного органа об отсутствии задолженности по налогам и обязательным платежам в бюджет, либо акт сверки",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Бизнес – план согласно Инструкции по составлению бизнес – плана инвестиционного проекта СПК (на бумажном и электронном носителях), паспорт проекта",
		Rus:            "Бизнес – план согласно Инструкции по составлению бизнес – плана инвестиционного проекта СПК (на бумажном и электронном носителях), паспорт проекта",
		Eng:            "Бизнес – план согласно Инструкции по составлению бизнес – плана инвестиционного проекта СПК (на бумажном и электронном носителях), паспорт проекта",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           1,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	//{
	//	IsAdditional:   false,
	//	Kaz:            "Документы, удостоверяющие право собственности на имущество, передаваемое в совместный с СПК проект (при наличии)",
	//	Rus:            "Документы, удостоверяющие право собственности на имущество, передаваемое в совместный с СПК проект (при наличии)",
	//	Eng:            "Документы, удостоверяющие право собственности на имущество, передаваемое в совместный с СПК проект (при наличии)",
	//	StartDate:      utils.GetCurrentTime(),
	//	DurationInDays: 3,
	//	GantaParentId:  0,
	//	
	//},
	//{
	//	IsAdditional:   false,
	//	Kaz:            "Справка об отсутствии обременения на имущество, передаваемое в совместный с СПК проект (при наличии)",
	//	Rus:            "Справка об отсутствии обременения на имущество, передаваемое в совместный с СПК проект (при наличии)",
	//	Eng:            "Справка об отсутствии обременения на имущество, передаваемое в совместный с СПК проект (при наличии)",
	//	StartDate:      utils.GetCurrentTime(),
	//	DurationInDays: 3,
	//	GantaParentId:  0,
	//	
	//},
	//{
	//	IsAdditional:   false,
	//	Kaz:            "Копия договоров займов (при наличии)",
	//	Rus:            "Копия договоров займов (при наличии)",
	//	Eng:            "Копия договоров займов (при наличии)",
	//	StartDate:      utils.GetCurrentTime(),
	//	DurationInDays: 3,
	//	GantaParentId:  0,
	//	
	//},
	//{
	//	IsAdditional:   false,
	//	Kaz:            "Форэскиз (по необходимости)",
	//	Rus:            "Форэскиз (по необходимости)",
	//	Eng:            "Форэскиз (по необходимости)",
	//	StartDate:      utils.GetCurrentTime(),
	//	DurationInDays: 3,
	//	GantaParentId:  0,
	//	
	//},
}

//var DefaultGantaChildren1LandDocs = []Ganta{
//	{
//		IsAdditional:   false,
//		Kaz:            "Заявки в ЦОН",
//		Rus:            "Заявки в ЦОН",
//		Eng:            "Заявки в ЦОН",
//		StartDate:      utils.GetCurrentTime(),
//		DurationInDays: 3,
//		GantaParentId:  0,
//		
//	},
//	{
//		IsAdditional:   false,
//		Kaz:            "ЦОН",
//		Rus:            "ЦОН",
//		Eng:            "ЦОН",
//		StartDate:      utils.GetCurrentTime(),
//		DurationInDays: 3,
//		GantaParentId:  0,
//		
//	},
//	{
//		IsAdditional:   false,
//		Kaz:            "ГУ Архитектуры Зем. комитет",
//		Rus:            "ГУ Архитектуры Зем. комитет",
//		Eng:            "ГУ Архитектуры Зем. комитет",
//		StartDate:      utils.GetCurrentTime(),
//		DurationInDays: 3,
//		
//	},
//	{
//		IsAdditional:   false,
//		Kaz:            "ГУ Архитектуры Зем. комитет",
//		Rus:            "ГУ Архитектуры Зем. комитет",
//		Eng:            "ГУ Архитектуры Зем. комитет",
//		StartDate:      utils.GetCurrentTime(),
//		DurationInDays: 3,
//		
//	},
//	{
//		IsAdditional:   false,
//		Kaz:            "Гос. Акт/ Выкуп права землепольз. на 5 лет",
//		Rus:            "Гос. Акт/ Выкуп права землепольз. на 5 лет",
//		Eng:            "Гос. Акт/ Выкуп права землепольз. на 5 лет",
//		StartDate:      utils.GetCurrentTime(),
//		DurationInDays: 3,
//		
//	},
//}
