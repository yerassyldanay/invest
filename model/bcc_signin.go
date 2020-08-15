package model

type SignIn struct {
	KeyUsername			string				`json:"key" validate:"nonzero"`
	Username			string				`json:"username" validate:"nonzero"`
	Password			string				`json:"password" validate:"nonzero"`
	Role				string				`json:"role" validate:"nonzero"`
	Id					uint64				`json:"id"`
}
