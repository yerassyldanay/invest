package templates

const (
	Base_email_notification_subject_kaz = "Ескерту! Сіз проектке қосылдыңыз"
	Base_email_notification_subject_rus = "Напоминание! Вас подключили к проекту"
	Base_email_notification_subject_eng = "Notification! You have been added to the project"

	Base_email_notification_page_kaz = `
	<div>
		Сіз келесі проектке қосылдыңыз!
		<b>Проект: </b> %s
		<div>
			<p>
				%s
			</p>
		</div>
	</div>
`
	Base_email_notification_page_rus = `
	<div>
		Вас подключили к следующему проекту!
		<b>Проект: </b> %s
		<div>
			<p>
				%s
			</p>
		</div>
	</div>
`
	Base_email_notification_page_eng = `
	<div>
		You have been added to the following project!
		<b>Project: </b> %s
		<div>
			<p>
				Description: %s
			</p>
		</div>
	</div>
`

	//Base_email_notification_html_kaz = ""
	//Base_email_notification_html_rus = ""
	//Base_email_notification_html_eng = ""
)
