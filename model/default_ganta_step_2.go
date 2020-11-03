package model

import (
	"invest/utils/constants"
	"invest/utils/helper"
)

var DefaultGantaParentsOfStep2 = []Ganta{
	{
		IsAdditional:   false,
		Kaz:            "Құжаттарды жүктеу",
		Rus:            "Загрузка документов инвестором",
		Eng:            "Uploading files by an investor",
		DurationInDays: 3,
		GantaChildren:  []Ganta{},

		Step:   2,
		Status: constants.ProjectStatusPendingInvestor,

		Responsible: constants.RoleInvestor,
		IsDone:      false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Құжаттарды тексеру",
		Rus:            "Проверка документов",
		Eng:            "Review of documents",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:   2,
		Status: constants.ProjectStatusPendingManager,

		Responsible: constants.RoleManager,
		IsDone:      false, // always false
		IsDocCheck:  true,
	},
	{
		IsAdditional:   false,
		Kaz:            "Инвестициялық комитет басшысының тексеруінде",
		Rus:            "Проверка главой инвестиционного комитета",
		Eng:            "Review by a head of an investment committee",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:   2,
		Status: constants.ProjectStatusPendingAdmin,

		Responsible: constants.RoleAdmin,
		IsDone:      false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Қаржы-қ / құқық-қ сараптама",
		Rus:            "Фин / юр эспертиза",
		Eng:            "A finance / jurisdiction expertise",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:        2,
		Status:      constants.ProjectStatusPendingExpert,
		Responsible: constants.RoleExpert,
		IsDocCheck:  true,
	},
	{
		IsAdditional:   false,
		Kaz:            "Инвестициялық комитет басшысының тексеруінде",
		Rus:            "Проверка главой инвестиционного комитета",
		Eng:            "Review by a head of an investment committee",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:   2,
		Status: constants.ProjectStatusPendingAdmin,

		Responsible: constants.RoleAdmin,
		IsDone:      false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "ӘКК Кеңесінің жобаны қарастыруы",
		Rus:            "На рассмотрение правления СПК",
		Eng:            "Consideration of the project by the Board",
		StartDate:      helper.GetCurrentTime(),
		DurationInDays: 21,
		GantaParentId:  0,

		Step:        2,
		Status:      constants.ProjectStatusPendingManager,
		Responsible: constants.RoleManager,

		GantaChildren:  []Ganta{},
	},
	{
		IsAdditional:   false,
		Kaz:            "Келісімге отыру",
		Rus:            "Заключение договора",
		Eng:            "Conclusion of the agreement",
		StartDate:      helper.GetCurrentTime(),
		DurationInDays: 15,
		GantaParentId:  0,

		Step:        2,
		Status:      constants.ProjectStatusPendingManager,
		Responsible: constants.RoleManager,

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
	StartDate:      helper.GetCurrentTime(),
	DurationInDays: 0,

	Step:        3,
	Status:      constants.ProjectStatusAgreement,
	Responsible: constants.RoleNobody,

	GantaChildren:  []Ganta{},
	NotToShow: true,
}
