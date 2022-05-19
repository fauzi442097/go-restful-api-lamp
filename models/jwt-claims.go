package models

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
}
