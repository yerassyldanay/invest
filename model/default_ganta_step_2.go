package model

import (
	"invest/utils"

)

var DefaultGantaParentsOfStep2 = []Ganta{
	{
		IsAdditional:   false,
		Kaz:            "Құжаттарды жүктеу [2]",
		Rus:            "Загрузка документов инвестором [2]",
		Eng:            "Uploading files by an investor [2]",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:           2,
		Status: 		utils.ProjectStatusPendingInvestor,

		Responsible:    utils.RoleInvestor,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Құжаттарды тексеру [2]",
		Rus:            "Проверка документов [2]",
		Eng:            "Review of documents [2]",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:           2,
		Status: 		utils.ProjectStatusPendingManager,

		Responsible:    utils.RoleManager,
		IsDone: 		false, // always false
		IsDocCheck: 	true,
	},
	{
		IsAdditional:   false,
		Kaz:            "Инвестициялық комитет басшысының тексеруі [2]",
		Rus:            "Проверка главой инвестиционного комитета [2]",
		Eng:            "Review by a head of an investment committee [2]",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:           2,
		Status: 		utils.ProjectStatusPendingAdmin,

		Responsible:    utils.RoleAdmin,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Қаржы-қ / құқық-қ сараптама [2]",
		Rus:            "Фин / юр эспертиза [2]",
		Eng:            "A finance / jurisdiction expertise [2]",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:           2,
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

		Step:           2,
		Status: 		utils.ProjectStatusPendingAdmin,

		Responsible:    utils.RoleAdmin,
		IsDone: 		false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "ӘКК Кеңесінің жобаны қарастыруы",
		Rus:            "Рассмотрение Проекта на Правлении СПК",
		Eng:            "Consideration of the project by the Board",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 21,
		GantaParentId:  0,

		Step:           2,
		Status: 		utils.ProjectStatusPendingManager,
		Responsible:    utils.RoleManager,

		GantaChildren:  []Ganta{},
	},
	{
		IsAdditional:   false,
		Kaz:            "Келісімге отыру",
		Rus:            "Заключение договора ДСД",
		Eng:            "Conclusion of the agreement",
		StartDate:      utils.GetCurrentTime(),
		DurationInDays: 15,
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

	DefaultGantaFinalStep,

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

var DefaultGantaFinalStep = Ganta{
	IsAdditional:   false,
	Kaz:            "Жоба барлық кезеңдерден өтті",
	Rus:            "Проект прошел все этапы",
	Eng:            "The project has passed all stages",
	StartDate:      utils.GetCurrentTime(),
	DurationInDays: 0,

	Step:           2,
	Status: 		utils.ProjectStatusAgreement,
	Responsible:    utils.RoleNobody,

	GantaChildren:  []Ganta{},
	NotToShow: true,
}
