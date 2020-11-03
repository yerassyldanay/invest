package model

import (
	"invest/utils/constants"
)

var DefaultDocuments1 = []Document{
	{
		Kaz:            "Заңды тұлғаны мемлекеттік тіркеу туралы куәлік",
		Rus:            "Свидетельство/справка о государственной регистрации юр. лица",
		Eng:            "Certificate of state registration of a legal entity",

		Step:        1,
		Responsible: constants.RoleInvestor, // this status will not be considered instead consider the status of the child
	},
	{
		Kaz:            "Жарғы, құрылтай жарғысы",
		Rus:            "Устав, учредительный договор",
		Eng:            "Articles of Association",

		Step:        1,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Заңды тұлғаның уәкілетті органының атқарушы органды тағайындау туралы шешімі",
		Rus:            "Решение Уполномоченного органа юридического лица о назначении исполнительного органа",
		Eng:            "Decision of the Authorized body of a legal entity on the appointment of an executive body",

		Step:        1,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Жарғылық капиталының 10%-дан астамы көлемі бар басшы, бас бухгалтер және құрылтайшы сертификатының көшірмелері",
		Rus:            "Копии удостоверения руководителя, главного бухгалтера и учредителей с более 10% уставного капитала",
		Eng:            "Copies of the certificate of the head, chief accountant and founders with more than 10% of the authorized capital",

		Step:        1,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Жоба іске асыруға қаражаттың бар екен растайтын құжаттар",
		Rus:            "Документы, подтверждающие наличие денежных средств на реализацию",
		Eng:            "Documents confirming the availability of fund for implementation",

		Step:        1,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Бірінші Несиелік Бюроның есебі",
		Rus:            "Отчет Первого Кредитного Бюро",
		Eng:            "Report of the First Credit Bureau",

		Step:        1,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Салық берешегінің және міндетті төлемдердің жоқтығы туралы анықтама немесе салыстыру актісі",
		Rus:            "Справка об отсутствии задолженности по налогам и обязательным платежам, либо акт сверки",
		Eng:            "Certificate of absence of tax arrears and obligatory payments, or reconciliation statement",

		Step:        1,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Нұсқаулыққа сәйкес бизнес жоспар, жоба паспорты",
		Rus:            "Бизнес – плана по инструкции, паспорт проекта",
		Eng:            "Business plan according to instructions, project passport",

		Step:        1,
		Responsible: constants.RoleInvestor,
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
