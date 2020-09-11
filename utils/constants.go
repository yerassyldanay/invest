package utils

const (
	GetLimitProjects = 20
)

const (
	RoleAdmin = "admin"
	RoleManager = "manager"
	RoleInvestor = "investor"
	RoleLawyer = "lawyer"
	RoleFinancier = "financier"
	/*
		this is needed to track what has been done by the system
	 */
	RoleSystem = "system"
)

const (
	ProjectStatusNewone = "newone"
	ProjectStatusInprogress = "inprogress"
	ProjectStatusDone = "done"
	ProjectStatusRejected = "rejected"

	ProjectStatusChangeTimeInHours = 48
)

const (
	FolderLogFiles = "logdir"
)

const (
	PanicLevel int = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

var NoNeedToAuth = []string {
	"/intest",
	"/api/check",
	"/v1/all/signup",
	"/v1/all/signin",
	"/v1/all/confirmation/email",
	"/v1/all/confirmation/phone",
	"/v1/all/organization",
}

var NoNeedToConfirmEmail = []string {}

const (
	KeyId = "Id"
	KeyRole = "Role"
	KeyRoleId = "Rid"
	KeyTime = "Time"
	KeyLanguage = "Lang"
	KeyOffset = "Offset"

	KeyEmailSubject = "subject"
	KeyEmailPlainText = "page"
	KeyEmailHtml = "html"

	DocTypeDocs = "docs"
	DocTypeComment = "comment"
)

const (
	AttemptToConnectToDb = 3
	TimeSecToSleepBetweenDbConn = 5

	MaxNumberOpenConnToDb = 5
	MaxNumberOfDigitsSentByEmail = 4
	MaxNumberOfCharactersSentByEmail = 30
)

const (
	ConstantDefaultNumberOfUsers = 3
)

const (
	CustomCostForHashing = 11
)

const (
	PositionAdmin = 0
	PositionInvestor = 1
	PositionManager = 2
	PositionLawyer = 3
	PositionFinancier = 4
	PositionSystem = 5
)

const (
	BaseEmailAddress = "yerassyl.danay@nu.edu.kz"
	BaseEmailName = "SPK"
)

const (
	HeaderAuth = "Authentication"
	HeaderContentLanguage = "Content-Language"
	HeaderAcceptLanguage = "Accept-Language"
	HeaderCustomStatus = "Custom-Status"
	HeaderContentType = "Content-Type"
	HeaderAuthorization = "Authorization"
)

const (
	CookieLanguageKeyWord = "lang"
)

const (
	DefaultContentLanguage = "kaz"
	ContentLanguageKk = "kaz"
	ContentLanguageRu = "rus"
	ContentLanguageEn = "eng"
	DefaultNotAllowedUserToDelete = 3
)

var LanguageMap = map[string]string{
	"kk": "kaz",
	"ru": "rus",
	"en": "eng",
}

const (
	AuthorizationAdminToken = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlX2lkIjoxLCJleHAiOiIyMDIwLTA4LTMwVDE3OjEzOjAzLjQwNDA2MDkxNyswNjowMCJ9.9FC_Ihr1gDLyJ8EX_wlwymECuOmKS8VeLCpW1RnO6WM`
)

const (
	GormSeqIdFinance = "finance"
)

const (
	RedisKeyForgetPassword = "forget_password"
)