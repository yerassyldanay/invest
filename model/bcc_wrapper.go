package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	UserId 			uint64			`json:"user_id"`
	RoleId			uint64			`json:"role_id"`
	RoleName		string			`json:"role_name"`
	Exp				time.Time		`json:"exp"`
	jwt.StandardClaims
}

type UserPermission struct {
	UserId				uint64				`json:"user_id"`
	Permission			string				`json:"permission"`
}

