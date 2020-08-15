package model

/*
	note: 2^32 = 4 294 967 296
*/
type User struct {
	Id				uint64				`json:"id" gorm:"AUTO_INCREMENT;primary key"`

	Username		string				`json:"username" validate:"required"`
	Password		string				`json:"password"`

	Fio				string				`json:"fio" validate:"required"`
	Position		string				`json:"position" validate:"required"`
	
	RoleId			uint64				`json:"role_id"`
	Role			Role				`json:"role" gorm:"foreignkey:RoleId"`

	EmailId			uint64				`json:"email_id"`
	Email			Email				`json:"email" gorm:"foreignkey:EmailId"`

	PhoneId			uint64				`json:"phone_id"`
	Phone			Phone				`json:"phone" gorm:"foreignkey:PhoneId"`

	Verified		bool				`json:"verified" gorm:"default:false"`
	Lang			string				`json:"-" gorm:"-"`

	Blocked			bool				`json:"blocked" gorm:"default:false"`
}

/*
	this returns the name of the table in the database
		gorm must automatically set the name by itself (adding 's' at the end)
		but it is worth to make sure that the name set correctly
*/
func (User) TableName() string {
	return "users"
}
