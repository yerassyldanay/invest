package model

import (
	"invest/utils/constants"
)

var DefaultDocuments2 = []Document{
	{
		Kaz:            "Акциялар шығарылымының проспектісі, акционерлердің тізілімінен үзінділер, акцияларды шығару нәтижелері туралы есеп, акцияларды орналастыру нәтижелері туралы есепті бекіту туралы хабарлама, бағалы қағаздар шығарылымын тіркеу туралы куәлік, тіркеушімен (тәуелсіз тіркеушімен) келісім - акционерлік қоғамдар үшін",
		Rus:            "Проспект выпуска акций, выписка из реестра акционеров, отчет об итогах выпуска эмиссии акций, уведомление об утверждении отчета об итогах размещения акций, свидетельство о регистрации выпуска ценных бумаг, договор с реестродержателем (независимым регистратором) - для акционерных обществ",
		Eng:            "Prospectus for the issue of shares, extract from the register of shareholders, report on the results of the issue of shares, notification of the approval of the report on the results of the placement of shares, certificate of registration of the issue of securities, agreement with the registrar (independent registrar) - for joint stock companies",

		Step:        2, // this status will not be considered instead consider the status of the child
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Жарғылық капиталды құруды растайтын құжаттар (шаруашылық серіктестік үшін)",
		Rus:            "Документы, подтверждающие формирование уставного капитала (для хозяйственного товарищества)",
		Eng:            "Documents confirming the formation of the authorized capital (for a business partnership)",

		Step:        2,
		Responsible: constants.RoleInvestor,
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
		Kaz:            "Қолтаңбалар мен мөр басылған үлгілері бар карточка (нотариалды куәландырылған)",
		Rus:            "Карточка с образцами подписей и оттиском печати (нотариально заверенная)",
		Eng:            "Card with samples of signatures and seal imprint (notarized)",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Компанияның соңғы 3 есепті жылдағы қаржылық есептілігінің көшірмесі, декларацияның көшірмесін қоса декларация көшірмелерін қоса бере отырып",
		Rus:            "Копия финансовой отчетности предприятия за последние 3 отчетных года с приложением копий налоговых деклараций с приложением расшифровки",
		Eng:            "A copy of the company's financial statements for the last 3 reporting years with the attachment of copies of tax returns with the attachment of a transcript",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Баланстың негізгі баптарына түсініктеме: дебиторлық және кредиторлық қарыз, пайда болу шарттары мен мерзімі, негізгі құралдар, шикізат, материалдар, дайын өнімдер көрсетілген дебиторлық және кредиторлық берешек",
		Rus:            "Расшифровка основных статей баланса: дебиторской и кредиторской задолженностей с указанием источника, условий и даты возникновения, основных средств, сырья, материалов, готовой продукции",
		Eng:            "Explanation of the main balance sheet items: accounts receivable and payable (indicating the source), conditions and date of occurrence, fixed assets, raw materials, materials, finished products",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},

	{
		Kaz:            "Қызмет көрсетуші банктің кезеңнің басындағы және аяғындағы ақшаның кіріс және шығыс қалдықтарын көрсете отырып, соңғы он екі айдағы айналымы туралы анықтамасы",
		Rus:            "Справка обслуживающего банка об оборотах за последние двенадцать месяцев, с указанием входящего и исходящего остатка денег на начало и конец периода",
		Eng:            "Certificate of a servicing bank about the turnover for the last twelve months, indicating the incoming and outgoing balance of money at the beginning and end of the period",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Банктердегі шоттар туралы ақпарат, несиелік берешектің бар / жоқтығы туралы ақпарат және картотека #2",
		Rus:            "Сведения о счетах в Банках, сведения о наличии/ отсутствии ссудной задолженности и картотеки #2",
		Eng:            "Information on bank accounts, information on cash / debt repayment and card files #2",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Бірінші несиелік бюроның есебі (ағымдағы күнге дейін жаңартылған)",
		Rus:            "Отчет Первого кредитного бюро (обновленный на текущую дату)",
		Eng:            "First Credit Bureau Report (updated to the current date)",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Салық берешегінің және бюджетке төленетін міндетті төлемдері жоқтығы туралы мемлекеттік органның анықтамасы немесе салыстыру актісі. (ағымдағы күнге дейін жаңартылған)",
		Rus:            "Справка из государственного органа об отсутствии задолженности по налогам и обязательным платежам в бюджет, либо акт сверки. (обновленную на текущую дату)",
		Eng:            "A certificate from a government agency about the absence of tax arrears and obligatory payments to the budget, or a reconciliation statement. (updated to the current date)",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Басқа құжаттарды (жарамды келісімшарттар, қаржылық құжаттар және т.б.) жоба менеджері жобаны қарау кезінде сұрайды",
		Rus:            "Другие документы (действующие контракты, финансовые документы и т.д.) запрашиваются менеджером проекта в ходе рассмотрения проекта",
		Eng:            "Other documents (valid contracts, financial documents, etc.) are requested by the project manager during the project review",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:         "ӘКК-нің инвестициялық жобасы бойынша бизнес-жоспарды құру жөніндегі нұсқаулыққа сәйкес бизнес-жоспар (қағаз және электрондық жеткізгіштерде), жоба паспорты (ағымдағы күнге дейін жаңартылған)",
		Rus:         "Бизнес – план согласно Инструкции по составлению бизнес – плана инвестиционного проекта СПК (на бумажном и электронном носителях), паспорт проекта (с корректировкой на текущую дату)",
		Eng:         "Business plan - in accordance with the Instruction - for an investment project of the SEC (hard or electronic version), project passport (with correction for the current date)",
		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "ӘКК-мен бірлескен жобаға берілген мүлікті бағалау туралы есеп",
		Rus:            "Отчет об оценке имущества, передаваемого в совместный с СПК проект",
		Eng:            "Report on the appraisal of property transferred to a joint project with the SEC",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
	{
		Kaz:            "Аффилиирленген компаниялар туралы ақпарат (акционер, бас директор)",
		Rus:            "Информация по аффилированным компаниям (акционера, первого руководителя)",
		Eng:            "Information on affiliated companies (shareholder, chief executive officer)",

		Step:        2,
		Responsible: constants.RoleInvestor,
	},
}
