package templates

const (
	Base_email_subject_kaz = "Қош келдіңіз"
	Base_email_subject_rus = "Добро пожаловать в СПК!"
	Base_email_subject_eng = "Welcome to SPK! Verify Code"
)

const (
	Base_email_html_kaz = `
		<div>
			<b>Тсс! Құпия сан:</b> %s
			<div>
				<p>
					Бұл сан 24 сағат бойы жарамды
				</p>
			</div>
			
			%s
		</div>
`
	Base_email_html_rus = `
		<div>
			<b>Тсс! Ваш секретный код:</b> %s
			<div>
				<p>
					Код валиден в течение 24 часа
				</p>
			</div>
			
			%s
		</div>
`
	Base_email_html_eng = `
		<div>
			<b>Your code is:</b> %s
			<div>
				<p>
					You code will be active for 24 hours
				</p>
			</div>

			%s
		</div>
`
)

const (
	Base_email_page_kaz = `
		<div>
			<b>Page! Тсс! Құпия сан:</b> %s
			<div>
				<p>
					Бұл сан 24 сағат бойы жарамды
				</p>
			</div>
			
			%s
		</div>
`
	Base_email_page_rus = `
		<div>
			<b>Page! Тсс! Ваш секретный код:</b> %s
			<div>
				<p>
					Код валиден в течение 24 часа
				</p>
			</div>
			
			%s
		</div>
`
	Base_email_page_eng = `
		<div>
			<b>Page! Tss! Your code is:</b> %s
			<div>
				<p>
					You code will be active for 24 hours
				</p>
			</div>

			%s
		</div>
`
)

const BaseUrlToConfirmEmail = "http://localhost:7000/v1/all/confirmation/email"
