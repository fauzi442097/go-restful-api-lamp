package services

import (
	"errors"
	"go-restful-api-lamp/config"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func GenerateToken(user models.User) (map[string]string, models.Claims) {
	// Create map token
	var tokenMap = map[string]string{}

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

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := t.SignedString([]byte(config.Secret))
	helper.PanicIfError(err)

	// Create Refresh Token
	rtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // exp 1 week
	})

	refreshToken, err := rtClaims.SignedString([]byte(config.Secret))
	if err != nil {
		helper.PanicIfError(err)
	}

	tokenMap["access_token"] = tokenString
	tokenMap["refresh_token"] = refreshToken

	return tokenMap, claims
}

func ValidateToken(tokenString string, userClaims *models.Claims) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			messageError := "unexpected signing method " + token.Header["alg"].(string)
			return nil, errors.New(messageError)
		}
		return []byte(config.Secret), nil
	})

	// claim := token.Claims.(jwt.MapClaims)
	// fmt.Println(claim["email"])
	// fmt.Println(claim["exp"])
	// fmt.Println(claim["iss"])
	// fmt.Println(claim["username"])

	return token, err
}
