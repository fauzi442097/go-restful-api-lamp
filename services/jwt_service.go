package services

import (
	"go-restful-api-lamp/config"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func GenerateToken(user models.User) (string, models.Claims) {
	jwtExpired, err := strconv.ParseInt(config.JwtExpiredAt, 10, 64)
	helper.PanicIfError(err)

	expirationTime := time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix()

	// create object claims
	claims := models.Claims{
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Issuer,
			ExpiresAt: jwt.NewTime(float64(expirationTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(config.Secret))
	helper.PanicIfError(err)

	return tokenString, claims
}

func ValidateToken(tokenString string, userClaims *models.Claims) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate alg
		// if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		// 	return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		// }
		return []byte(config.Secret), nil
	})

	return token, err
}
