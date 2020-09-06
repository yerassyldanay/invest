package templates

var Base_message_map_2_forget_password = map[string]map[string]string{
	"subject": {
		"kaz": "Құпия сөзді жаңарту",
		"rus": "Пароль сборшен",
		"eng": "Password was renewed",
	},

	"html": {
		"kaz": `	<div>
				Құпия сөзді жаңарту
				<div>
					<p>
						%s
					</p>
				</div>
			</div>
		`,
		"rus": `
			<div>
				Пароль сборшен
				<div>
					<p>
						%s
					</p>
				</div>
			</div>
		`,
		"eng": `
			<div>
				Password was renewed
				<b>Project: </b> %s
				<div>
					<p>
						Description: %s
					</p>
				</div>
			</div>
		`,
	},
}