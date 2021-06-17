package model

import (
	"github.com/yerassyldanay/invest/utils/constants"
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

		Step:   1,
		Status: constants.ProjectStatusPendingInvestor,

		Responsible: constants.RoleInvestor,
		IsDone:      false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Менеджерді тағайындау",
		Rus:            "Назначение менеджера",
		Eng:            "Assignment of a manager",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:   1,
		Status: constants.ProjectStatusPendingAdmin,

		Responsible: constants.RoleAdmin,
		IsDone:      false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "Құжаттарды тексеру",
		Rus:            "Проверка документов",
		Eng:            "Review of documents",
		DurationInDays: 2,
		GantaChildren:  []Ganta{},

		Step:   1,
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

		Step:   1,
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

		Step:        1,
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

		Step:   1,
		Status: constants.ProjectStatusPendingAdmin,

		Responsible: constants.RoleAdmin,
		IsDone:      false, // always false
	},
	{
		IsAdditional:   false,
		Kaz:            "ИК-тің қарауында",
		Rus:            "Рассмотрение на ИК СПК",
		Eng:            "Review by an invest committee",
		DurationInDays: 6,
		GantaChildren:  []Ganta{},

		Step:        1,
		Status:      constants.ProjectStatusPendingInvCommittee,
		Responsible: constants.RoleManager,
	},
	{
		IsAdditional:   false,
		Kaz:            "Жерді пайдалану құқығына құжаттарды тіркеу",
		Rus:            "Оформление документов на право землепользования",
		Eng:            "Registration of documents for the right to land use",
		DurationInDays: 9 * 30,
		GantaChildren:  []Ganta{},

		Step:        1,
		Status:      constants.ProjectStatusPendingManager,
		Responsible: constants.RoleManager,
	},
}
