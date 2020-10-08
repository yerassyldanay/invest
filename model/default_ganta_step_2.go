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
		
		GantaChildren:  []Ganta{},
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
		Eng:            "Протокол Решения Правления",
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

var DefaultGantaFinalStep = Ganta{
	IsAdditional:   false,
	Kaz:            "Проект прошел все этапы",
	Rus:            "Проект прошел все этапы",
	Eng:            "Проект прошел все этапы",
	StartDate:      utils.GetCurrentTime(),
	DurationInDays: 0,

	Step:           3,
	Status: 		utils.ProjectStatusAgreement,
	Responsible:    utils.RoleNobody,

	GantaChildren:  []Ganta{},
}
