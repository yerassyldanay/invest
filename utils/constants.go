package utils

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
	"/test",
	"/api/check",
	"/v1/all/signup",
	"/v1/all/signin",
	"/v1/all/confirmation/email",
	"/v1/all/confirmation/phone",
}

const (
	KeyId = "id"
	KeyRole = "role"
	KeyRoleId = "rid"
	KeyTime = "time"
	KeyLanguage = "lang"
	KeyOffset = "offset"

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
)

const (
	DefaultLContentanguage = "kaz"
	DefaultNotAllowedUserToDelete = 3
)

const (
	ProjectStatusNotConfirmed = "not confirmed"
	ProjectStatusConfirmed	= "confirmed"
	ProjectStatusReturnedToChange = "returned for revision"
	ProjectStatusBlocked = "blocked"

	ProjectStatusChangeTimeInHours = 48
)