package utils

var MapRole = map[string]map[string]string {
	RoleAdmin: {
		"kaz": "админ",
		"rus": "админ",
		"eng": "admin",
	},
	RoleManager: {
		"kaz": "менеджер",
		"rus": "менеджер",
		"eng": "manager",
	},
	RoleExpert: {
		"kaz": "эксперт",
		"rus": "эксперт",
		"eng": "expert",
	},
	RoleNobody: {
		"kaz": "ешкім",
		"rus": "никто",
		"eng": "nobody",
	},
	RoleSpk: {
		"kaz": "ӘКК",
		"rus": "СПК",
		"eng": "SEC",
	},
}

const (
	RoleInvestor = "investor"
	RoleManager = "manager"
	RoleExpert = "expert"
	RoleAdmin = "admin"
	RoleSpk = "spk"
	RoleNobody = "nobody"
)

const (
	ProjectStatusPendingInvestor = "pending_investor"
	ProjectStatusPendingManager = "pending_manager"
	ProjectStatusPendingAdmin = "pending_admin"
	ProjectStatusPendingExpert = "pending_expert"
	ProjectStatusPendingInvCommittee = "invest_committee"
	ProjectStatusRegistrationOfLandPlot = "reg_land_plot"
	ProjectStatusPendingBoard = "pending_board"
	ProjectStatusAgreement = "agreement"

	ProjectStatusDelay = "delay"
	ProjectStatusReject = "reject"
	ProjectStatusReconsider = "reconsider"
	ProjectStatusAccept = "accept"

	ProjectStatusNewOne = "new_one"

	ProjectStatusChangeTimeInHours = 48
)

var MapProjectStatusFirstStatusThenLang = map[string]map[string]string{
	ProjectStatusPendingInvestor: {
		"kaz": "Инвестордың қарауында",
		"rus": "На рассмотрении инвестора",
		"eng": "Review by an investor",
	},
	ProjectStatusPendingManager: {
		"kaz": "Менеджердің қарауында",
		"rus": "На рассмотрении менеджера",
		"eng": "Review by a manager",
	},
	ProjectStatusPendingAdmin : {
		"kaz": "Админнің қараунда",
		"rus": "На рассмотрении админа",
		"eng": "Review by an admin",
	},
	ProjectStatusPendingExpert : {
		"kaz": "Эксперттің қарауында",
		"rus": "На рассмотрении эксперта",
		"eng": "Review by an expert",
	},
	ProjectStatusPendingInvCommittee : {
		"kaz": "Инвест коммитеттің қарауында",
		"rus": "На рассмотрении инвестиционного коммитета",
		"eng": "Review by an Invest Committee",
	},
	ProjectStatusRegistrationOfLandPlot : {
		"kaz": "Жер учаскесін тіркеу",
		"rus": "На оформлении земельного участка",
		"eng": "Registration of land plot",
	},
	ProjectStatusPendingBoard : {
		"kaz": "ӘКК кеңесінің қарауында",
		"rus": "На рассмотрении правления СПК",
		"eng": "Review by the Board of SEC",
	},
	ProjectStatusAgreement : {
		"kaz": "Жоба барлық кезеңдерден өтті",
		"rus": "Проек прошел все этапы",
		"eng": "The project has passed all stages",
	},
	ProjectStatusDelay : {
		"kaz": "Процесс кідірісте",
		"rus": "Задержка",
		"eng": "Delay",
	},
	ProjectStatusReject : {
		"kaz": "Қабылданбады",
		"rus": "Отклонен",
		"eng": "Rejected",
	},
	ProjectStatusReconsider : {
		"kaz": "Қайта қарау",
		"rus": "На доработке",
		"eng": "Reconsideration",
	},
	ProjectStatusAccept : {
		"kaz": "Қабылданды",
		"rus": "Принято",
		"eng": "Accepted",
	},
	ProjectStatusNewOne : {
		"kaz": "Жаңа",
		"rus": "Новый",
		"eng": "New",
	},
}
