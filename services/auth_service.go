package services

import (
	"go-restful-api-lamp/dto"
	"go-restful-api-lamp/exception"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/models"
	"go-restful-api-lamp/repositories"
	"go-restful-api-lamp/utils/transaction"

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

	// Generate Token JWT
	token, claims := GenerateToken(user)

	// parsing to dto.user response
	userResponse := dto.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	helper.PanicIfError(err)

	dataResponse := dto.LoginResponse{
		User:         userResponse,
		Token:        token,
		TokenExpired: claims.StandardClaims.ExpiresAt,
	}

	return dataResponse
}
