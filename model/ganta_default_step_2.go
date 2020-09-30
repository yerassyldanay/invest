package model

import (
	"invest/utils"

)

var DefaultGantaParentsOfStep2 = []Ganta{
	{
		IsAdditional:   false,
		Kaz:            "Пакет документов предоставляемый Инициатором проекта на втором этапе",
		Rus:            "Пакет документов предоставляемый Инициатором проекта на втором этапе",
		Eng:            "Пакет документов предоставляемый Инициатором проекта на втором этапе",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 10,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
		
		GantaChildren:  DefaultGantaChildren2,
	},
	{
		IsAdditional:   false,
		Kaz:            "Экспертиза Проекта (ДИП, ДЭП, ЮД, ДБиНУ)",
		Rus:            "Экспертиза Проекта (ДИП, ДЭП, ЮД, ДБиНУ)",
		Eng:            "Экспертиза Проекта (ДИП, ДЭП, ЮД, ДБиНУ)",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 10,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingExpert,
		Responsible:    utils.RoleExpert,

		GantaChildren:  []Ganta{},
		IsDocCheck: 	true,
	},
	{
		IsAdditional:   false,
		Kaz:            "Пояснительная записка на Правление СПК с приложением пакета документов",
		Rus:            "Пояснительная записка на Правление СПК с приложением пакета документов",
		Eng:            "Пояснительная записка на Правление СПК с приложением пакета документов",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 10,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingManager,
		Responsible:    utils.RoleManager,

		GantaChildren:  []Ganta{},
	},
	{
		IsAdditional:   false,
		Kaz:            "Рассмотрение Проекта на Правлении СПК",
		Rus:            "Рассмотрение Проекта на Правлении СПК",
		Eng:            "Рассмотрение Проекта на Правлении СПК",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 10,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingManager,
		Responsible:    utils.RoleManager,

		GantaChildren:  []Ganta{},
	},
	{
		IsAdditional:   false,
		Kaz:            "Протокол Решения Правления",
		Rus:            "Протокол Решения Правления",
		Eng:            "Протокол Решения Правления)",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 10,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingManager,
		Responsible:    utils.RoleManager,

		GantaChildren:  []Ganta{},
	},
	{
		IsAdditional:   false,
		Kaz:            "Заключение договора ДСД / или создание СП",
		Rus:            "Заключение договора ДСД / или создание СП",
		Eng:            "Заключение договора ДСД / или создание СП",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 10,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingManager,
		Responsible:    utils.RoleManager,

		GantaChildren: []Ganta{
			//{
			//	IsAdditional:   false,
			//	Kaz:            "Исполнение условий договора ДСД/СП",
			//	Rus:            "Исполнение условий договора ДСД/СП",
			//	Eng:            "Исполнение условий договора ДСД/СП",
			//	StartDate:      utils.GetCurrentTime(),
			//	DurationInDays: 10,
			//	GantaParentId:  0,
			//	
			//},
		},
	},
	//{
	//	IsAdditional:   false,
	//	Kaz:            "Переход на инвестиционную стадию",
	//	Rus:            "Переход на инвестиционную стадию",
	//	Eng:            "Переход на инвестиционную стадию",
	//	StartDate:      utils.GetCurrentTime(),
	//	DurationInDays: 10,
	//	GantaParentId:  0,
	//	
	//	GantaChildren:  []Ganta{},
	//},
}

var DefaultGantaChildren2 = []Ganta{
	{
		IsAdditional:   false,
		Kaz:            "Проспект выпуска акций, Выписка из реестра акционеров, Отчет об итогах выпуска эмиссии акций, уведомление об утверждении отчета об итогах размещения акций, Свидетельство о регистрации выпуска ценных бумаг, Договор с реестродержателем (независимым регистратором) - для акционерных обществ",
		Rus:            "Проспект выпуска акций, Выписка из реестра акционеров, Отчет об итогах выпуска эмиссии акций, уведомление об утверждении отчета об итогах размещения акций, Свидетельство о регистрации выпуска ценных бумаг, Договор с реестродержателем (независимым регистратором) - для акционерных обществ",
		Eng:            "Проспект выпуска акций, Выписка из реестра акционеров, Отчет об итогах выпуска эмиссии акций, уведомление об утверждении отчета об итогах размещения акций, Свидетельство о регистрации выпуска ценных бумаг, Договор с реестродержателем (независимым регистратором) - для акционерных обществ",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,                                 // will be set

		Step:           2, // this status will not be considered instead consider the status of the child
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Документы, подтверждающие формирование уставного капитала (для хозяйственного товарищества)",
		Rus:            "Документы, подтверждающие формирование уставного капитала (для хозяйственного товарищества)",
		Eng:            "Документы, подтверждающие формирование уставного капитала (для хозяйственного товарищества)",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	//{
	//	IsAdditional:   false,
	//	Kaz:            "Лицензия – если вид деятельности лицензируемый",
	//	Rus:            "Лицензия – если вид деятельности лицензируемый",
	//	Eng:            "Лицензия – если вид деятельности лицензируемый",
	//	StartDate:      utils.GetCurrentTime(),
	//	DurationInDays: 3,
	//	GantaParentId:  0,
	//	Step:           2,
	//	
	//},
	{
		IsAdditional:   false,
		Kaz:            "Карточка с образцами подписей и оттиском печати (нотариально заверенная)",
		Rus:            "Карточка с образцами подписей и оттиском печати (нотариально заверенная)",
		Eng:            "Карточка с образцами подписей и оттиском печати (нотариально заверенная)",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Копия финансовой отчетности предприятия за последние 3 отчетных года с приложением копий налоговых деклараций с приложением расшифровки",
		Rus:            "Копия финансовой отчетности предприятия за последние 3 отчетных года с приложением копий налоговых деклараций с приложением расшифровки",
		Eng:            "Копия финансовой отчетности предприятия за последние 3 отчетных года с приложением копий налоговых деклараций с приложением расшифровки",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Расшифровка основных статей баланса: дебиторской и кредиторской задолженностей с указанием источника, условий и даты возникновения, основных средств, сырья, материалов, готовой продукции",
		Rus:            "Расшифровка основных статей баланса: дебиторской и кредиторской задолженностей с указанием источника, условий и даты возникновения, основных средств, сырья, материалов, готовой продукции",
		Eng:            "Расшифровка основных статей баланса: дебиторской и кредиторской задолженностей с указанием источника, условий и даты возникновения, основных средств, сырья, материалов, готовой продукции",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},

	{
		IsAdditional:   false,
		Kaz:            "Справка обслуживающего банка об оборотах за последние двенадцать месяцев, с указанием входящего и исходящего остатка денег на начало и конец периода",
		Rus:            "Справка обслуживающего банка об оборотах за последние двенадцать месяцев, с указанием входящего и исходящего остатка денег на начало и конец периода",
		Eng:            "Справка обслуживающего банка об оборотах за последние двенадцать месяцев, с указанием входящего и исходящего остатка денег на начало и конец периода",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Сведения о счетах в Банках, сведения о наличии/ отсутствии ссудной задолженности и картотеки No2",
		Rus:            "Сведения о счетах в Банках, сведения о наличии/ отсутствии ссудной задолженности и картотеки No2",
		Eng:            "Сведения о счетах в Банках, сведения о наличии/ отсутствии ссудной задолженности и картотеки No2",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Отчет Первого кредитного бюро (обновленный, на текущую дату)",
		Rus:            "Отчет Первого кредитного бюро (обновленный, на текущую дату)",
		Eng:            "Отчет Первого кредитного бюро (обновленный, на текущую дату)",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Справка из государственного органа об отсутствии задолженности по налогам и обязательным платежам в бюджет, либо акт сверки. (обновленную на текущую дату)",
		Rus:            "Справка из государственного органа об отсутствии задолженности по налогам и обязательным платежам в бюджет, либо акт сверки. (обновленную на текущую дату)",
		Eng:            "Справка из государственного органа об отсутствии задолженности по налогам и обязательным платежам в бюджет, либо акт сверки. (обновленную на текущую дату)",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Другие документы (действующие контракты, финансовые документы и т.д.) запрашиваются менеджером проекта в ходе рассмотрения проекта",
		Rus:            "Другие документы (действующие контракты, финансовые документы и т.д.) запрашиваются менеджером проекта в ходе рассмотрения проекта",
		Eng:            "Другие документы (действующие контракты, финансовые документы и т.д.) запрашиваются менеджером проекта в ходе рассмотрения проекта",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:  false,
		Kaz:           "Бизнес – план согласно Инструкции по составлению бизнес – плана инвестиционного проекта СПК (на бумажном и электронном носителях), паспорт проекта (с корректировкой на текущую дату)",
		Rus:           "Бизнес – план согласно Инструкции по составлению бизнес – плана инвестиционного проекта СПК (на бумажном и электронном носителях), паспорт проекта (с корректировкой на текущую дату)",
		Eng:           "Бизнес – план согласно Инструкции по составлению бизнес – плана инвестиционного проекта СПК (на бумажном и электронном носителях), паспорт проекта (с корректировкой на текущую дату)",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Отчет об оценке имущества, передаваемого в совместный с СПК проект",
		Rus:            "Отчет об оценке имущества, передаваемого в совместный с СПК проект",
		Eng:            "Отчет об оценке имущества, передаваемого в совместный с СПК проект",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
	{
		IsAdditional:   false,
		Kaz:            "Информация по аффилированным компаниям (акционера, первого руководителя)",
		Rus:            "Информация по аффилированным компаниям (акционера, первого руководителя)",
		Eng:            "Информация по аффилированным компаниям (акционера, первого руководителя)",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 3,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,
		Responsible:    utils.RoleInvestor,
	},
}
