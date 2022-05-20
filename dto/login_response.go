package dto

import "github.com/dgrijalva/jwt-go/v4"

type LoginResponse struct {
	User         UserResponse `json:"user"`
	Token        string       `json:"token"`
	TokenExpired *jwt.Time    `json:"token_expired"`
}
