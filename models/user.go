package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserTokenClaims struct {
	jwt.RegisteredClaims
	Id int `json:"id"`
}
