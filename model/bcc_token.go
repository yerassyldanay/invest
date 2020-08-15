package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	UserId 			uint64			`json:"user_id"`
	RoleId			uint64			`json:"role_id"`
	Exp				time.Time		`json:"exp"`
	jwt.StandardClaims
}

