package templates

var Base_message_map_1_welcome = map[string]map[string]string{
	"subject": map[string]string{
		"eng": "Welcome to SPK! Verify Code",
		"rus": "Добро пожаловать в СПК!",
		"kaz": "Қош келдіңіз",
	},
	"html": map[string]string{
		"eng": `
				<div>
					<b>Тсс! Құпия сан:</b> %s
					<div>
						<p>
							Бұл сан 24 сағат бойы жарамды
						</p>
					</div>
					
					%s
				</div>
		`,
		"rus": `
				<div>
					<b>Тсс! Ваш секретный код:</b> %s
					<div>
						<p>
							Код валиден в течение 24 часа
						</p>
					</div>
					
					%s
				</div>
		`,
		"kaz": `
				<div>
					<b>Your code is:</b> %s
					<div>
						<p>
							You code will be active for 24 hours
						</p>
					</div>
		
					%s
				</div>
		`,
	},
	"page": map[string]string{
		"eng": ``,
		"rus": ``,
		"kaz": ``,
	},
}
