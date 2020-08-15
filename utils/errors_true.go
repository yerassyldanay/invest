package utils

var ErrorInvalidPassword = map[string]interface{}{
	"eng": "invalid password",
	"rus": "невалидный пароль",
	"kaz": "құпия сөз қате",
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

var ErrorAccountNotVerified = map[string]interface{}{
	"eng": "account is not verified",
	"rus": "подтвердите аккаунт",
	"kaz": "аккаунт расталмаған",
}

var ErrorCouldNotGet = map[string]interface{}{
	"eng": "failed to get all civil servants",
	"rus": "",
	"kaz": "",
}

var ErrorMethodNotAllowed = map[string]interface{}{
	"eng": "this functionality is not allowed for you",
	"rus": "",
	"kaz": "",
}

var ErrorInternalServerError = map[string]interface{} {
	"eng": "an internal error has occurred",
	"rus": "произошла внутренняя ошибка",
	"kaz": "серверлік қате пайда болды",
}

var ErrorInternalDbError = map[string]interface{} {
	"eng": "failed create / update / remove a row in db",
	"rus": "не удалось создать / обновить / удалить строку в бд",
	"kaz": "мәліметті құру / жаңарту / жою сәтсіз аяқталды",
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

