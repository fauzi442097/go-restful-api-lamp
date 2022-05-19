package services

import (
	"go-restful-api-lamp/config"
	"go-restful-api-lamp/dto"
	"go-restful-api-lamp/exception"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/models"
	"go-restful-api-lamp/repositories"
	"go-restful-api-lamp/utils/transaction"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(user dto.LoginRequest) dto.LoginResponse
}

type AuthServiceImpl struct {
	db         *gorm.DB
	repository repositories.AuthRepository
	Validate   *validator.Validate
}

// NewAuthService returns new NewAuthService.
func NewAuthService(repo repositories.AuthRepository, db *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		repository: repo,
		db:         db,
		Validate:   validate,
	}
}

func (service *AuthServiceImpl) Login(userRequest dto.LoginRequest) dto.LoginResponse {

	credential := userRequest

	// Validate
	err := service.Validate.Struct(credential)
	helper.PanicIfError(err)

	// Start Transaction
	tx := service.db.Begin()
	defer transaction.CommitOrRollback(tx)

	userModel := models.User{}
	err = smapping.FillStruct(&userModel, smapping.MapFields(&credential))
	helper.PanicIfError(err)

	user, err := service.repository.Login(tx, userModel)
	if err != nil {
		panic(exception.NewErrorUnauthenticated("Username atau Password salah"))
	}

	hashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(userRequest.Password))
	if err != nil {
		// Password do not match
		panic(exception.NewErrorUnauthenticated("Username atau Password salah"))
	}

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

	// parsing to dto.user response
	userResponse := dto.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	helper.PanicIfError(err)

	dataResponse := dto.LoginResponse{
		User:  userResponse,
		Token: tokenString,
	}

	return dataResponse

}
