package model

type SignIn struct {
	KeyUsername			string				`json:"key"`
	Value				string				`json:"value"`
	Password			string				`json:"password"`
	Id					uint64				`json:"id"`
	TokenCompound				string				`json:"-,omitempty"`
}
