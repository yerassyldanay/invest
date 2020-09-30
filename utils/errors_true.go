package utils

var ErrorInvalidPassword = map[string]interface{}{
	"eng": "invalid password: password must contain 8 to 20 characters (upper, lower letters & digits)",
	"rus": "невалидный пароль: пароль должен состоять из от 8 до 20 знаков (большие и маленькие буквы & цифры)",
	"kaz": "құпия сөз қате: құпия сөз 8-20 таңбадан тұруы (үлкен және кіші әріптердің санның болу) қажет",
}

var ErrorInvalidParameters = map[string]interface{}{
	"eng": "invalid parameters have been passed",
	"rus": "невалидные параметры переданы",
	"kaz": "қате параметрлер берілген",
}

var ErrorNoSuchUser = map[string]interface{}{
	"eng": "there is no such user",
	"rus": "пользователь не существует",
	"kaz": "қолданушы тіркелмеген",
}

var ErrorNotFound = map[string]interface{} {
	"eng": "not found",
	"rus": "",
	"kaz": "",
}

var ErrorAccountNotVerified = map[string]interface{}{
	"eng": "account is not verified",
	"rus": "подтвердите аккаунт",
	"kaz": "аккаунт расталмаған",
}

var ErrorMethodNotAllowed = map[string]interface{}{
	"eng": "this functionality is not allowed for you",
	"rus": "вы не можете использовать эту функцию",
	"kaz": "сіз бұл функцияны қолдана алмайсыз",
}

var ErrorInternalServerError = map[string]interface{} {
	"eng": "an internal error has occurred",
	"rus": "произошла внутренняя ошибка",
	"kaz": "серверлік қате пайда болды",
}

var ErrorInternalDbError = map[string]interface{} {
	"eng": "failed create / update / remove / get a row from db",
	"rus": "не удалось создать / обновить / удалить / получить строку в бд",
	"kaz": "мәліметті құру / жаңарту / жою / алу сәтсіз аяқталды",
}

var ErrorEmailIsNotVerified = map[string]interface{} {
	"eng": "the account is not confirmed | please, confirm email address",
	"rus": "аккаунт не подтвержден | пожалуйста, подтвердите адрес электронной почты",
	"kaz": "аккаунт расталмаған | электрондық пошта мекенжайы арқылы растаңыз",
}

var ErrorPhoneNumberIsNotVerified = map[string]interface{} {
	"eng": "please, confirm the phone number",
	"rus": "пожалуйста, подтвердите номер телефона",
	"kaz": "телефон нөміріңізді растаңыз",
}

var ErrorFailedToMakeChanges = map[string]interface{} {
	"eng": "failed to make changes to the db",
}

var ErrorInternalIssueOrInvalidPassword = map[string]interface{}{
	"eng": "internal problem has occurred or the password is invalid",
	"rus": "произошла внутренняя ошибка или неверный пароль",
	"kaz": "сервисте қателік орын алды немесе құпия сөз жарамсыз",
}

var ErrorEmailIsAreadyInUse = map[string]interface{}{
	"eng": "email is already in use",
	"rus": "электронная почта уже используется",
	"kaz": "электрондық почта қолданыста",
}

var ErrorUsernameOrFioIsAreadyInUse = map[string]interface{}{
	"eng": "username or fio is already in use",
	"rus": "логин или фио уже используется",
	"kaz": "логин немесе адам есімі қолданыста",
}

var ErrorFailedToUpdateSomeValues = map[string]interface{}{
	"eng": "failed to update some values",
	"rus": "не удалось обновить некоторые значения",
	"kaz": "кейбір мәндерді жаңарта алмады",
}

var ErrorExternalServiceErrorNoOrganizationInfo = map[string]interface{} {
	"eng": "failed to obtain organization info",
}

var ErrorCouldNotSendEmail = map[string]interface{} {
	"eng": "failed to send an email",
}

var ErrorAlreadySentLinkToEmail = map[string]interface{}{
	"eng": "please, check your email. A link has been sent to your email address | Please, use this password",
	"rus": "пожалуйста, проверьте почту. Ссылка была отправлена на ваш электронный адрес | Пожалуйста, используйте этот пароль",
	"kaz": "почтаңызды тексеруді өтінеміз. Электрондық поштаңызға сілтеме жіберілді | Осы құпия сөзді қолданыңыз",
}

var ErrorFailedToCreateAnAccount = map[string]interface{}{
	"eng": "failed to create an account",
	"rus": "не удалось создать аккаунт",
	"kaz": "аккаунт ашу мүмкін болмады",
}

var ErrorDupicateKeyOnDb = map[string]interface{} {
	"eng": "failed to create, probably, because some fields are repeated",
	"rus": "не удалось создать, возможно, потому что некоторые поля повторяются",
	"kaz": "енгізілген кейбір ақпарат қайталанады",
}

var ErrorTokenInvalidOrExpired = map[string]interface{}{
	"eng": "invalid or expired token",
	"rus": "токен невалидный или устарел",
	"kaz": "токен дүрыс емес немесе ескірген",
}

