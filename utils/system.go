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
		"kaz": "",
		"rus": "",
		"eng": "",
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
		"kaz": "На рассмотрении инвестора",
		"rus": "На рассмотрении инвестора",
		"eng": "На рассмотрении инвестора",
	},
	ProjectStatusPendingManager: {
		"kaz": "На рассмотрении менеджера",
		"rus": "На рассмотрении менеджера",
		"eng": "На рассмотрении менеджера",
	},
	ProjectStatusPendingAdmin : {
		"kaz": "На рассмотрении админа",
		"rus": "На рассмотрении админа",
		"eng": "На рассмотрении админа",
	},
	ProjectStatusPendingExpert : {
		"kaz": "На рассмотрении эксперта",
		"rus": "На рассмотрении эксперта",
		"eng": "На рассмотрении эксперта",
	},
	ProjectStatusPendingInvCommittee : {
		"kaz": "На рассмотрении инвестиционного коммитета",
		"rus": "На рассмотрении инвестиционного коммитета",
		"eng": "На рассмотрении инвестиционного коммитета",
	},
	ProjectStatusRegistrationOfLandPlot : {
		"kaz": "На оформлении земельного участка",
		"rus": "На оформлении земельного участка",
		"eng": "На оформлении земельного участка",
	},
	ProjectStatusPendingBoard : {
		"kaz": "На рассмотрении правления СПК",
		"rus": "На рассмотрении правления СПК",
		"eng": "На рассмотрении правления СПК",
	},
	ProjectStatusAgreement : {
		"kaz": "Проек прошел все этапы",
		"rus": "Проек прошел все этапы",
		"eng": "Проек прошел все этапыа",
	},
	ProjectStatusDelay : {
		"kaz": "Задержка",
		"rus": "Задержка",
		"eng": "Задержка",
	},
	ProjectStatusReject : {
		"kaz": "Отклонен",
		"rus": "Отклонен",
		"eng": "Отклонен",
	},
	ProjectStatusReconsider : {
		"kaz": "На доработке",
		"rus": "На доработке",
		"eng": "На доработке",
	},
	ProjectStatusAccept : {
		"kaz": "Принято",
		"rus": "Принято",
		"eng": "Принято",
	},
	ProjectStatusNewOne : {
		"kaz": "Новый",
		"rus": "Новый",
		"eng": "Новый",
	},
}
