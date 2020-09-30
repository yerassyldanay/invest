package model

import "github.com/jinzhu/gorm"

func Update_sequence_id_thus_avoid_duplicate_primary_key_error(tx *gorm.DB, bycase string) (err error) {
	/*

	 */
	var main_query string

	/*
		Solution to ERROR:  duplicate key violates unique constraint
	*/
	switch bycase {
	case "costs":
		main_query = `select setval('costs_id_seq', (select max(id) from costs) + 1);`
	case "finances":
		main_query = `select setval('finances_id_seq', (select max(id) from finances) + 1);`
	case "gantas":
		main_query = `select setval('gantas_id_seq', (select coalesce(max(id), 0) as id from gantas) + 1);`
	case "emails":
		main_query = `select setval('emails_id_seq', (select max(id) from emails) + 1);`
	case "phones":
		main_query = `select setval('phones_id_seq', (select max(id) from phones) + 1);`
	case "users":
		main_query = `select setval('users_id_seq', (select max(id) from users) + 1);`
	case "roles":
		main_query = `select setval('roles_id_seq', (select max(id) from roles) + 1);`
	default:
		main_query = `select current_time;`
	}

	err = tx.Exec(main_query).Error
	return err
}
