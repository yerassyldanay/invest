package model

import "github.com/jinzhu/gorm"

func Update_sequence_id_thus_avoid_duplicate_primary_key_error(tx *gorm.DB, bycase string) error {
	/*

	 */
	var main_query string

	/*
		Solution to ERROR:  duplicate key violates unique constraint
	*/
	switch bycase {
	case "finance":
		main_query = `
			select setval('finresults_id_seq', (select max(id) from finresults) + 1);
			select setval('finances_id_seq', (select max(id) from finances) + 1);
		`
	default:
		main_query = `
			select setval('emails_id_seq', (select max(id) from emails) + 1);
			select setval('phones_id_seq', (select max(id) from phones) + 1);
			select setval('users_id_seq', (select max(id) from users) + 1);
			select setval('roles_id_seq', (select max(id) from roles) + 1);
		`
	}

	return tx.Exec(main_query).Error
}
