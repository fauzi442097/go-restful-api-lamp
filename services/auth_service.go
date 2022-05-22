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
	Register(userRegister dto.RegisterRequest) error
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
	tokenMap, _ := GenerateToken(user)

	// parsing to dto.user response
	userResponse := dto.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	helper.PanicIfError(err)

	dataResponse := dto.LoginResponse{
		User:         userResponse,
		Token:        tokenMap["access_token"],
		RefreshToken: tokenMap["refresh_token"],
	}

	return dataResponse
}

func (service *AuthServiceImpl) Register(userRegister dto.RegisterRequest) error {

	// Validasi
	err := service.Validate.Struct(&userRegister)
	helper.PanicIfError(err)

	// Start Trasaction
	tx := service.db.Begin()
	defer transaction.CommitOrRollback(tx)

	// Get Last User
	lastUserModel := models.User{}
	lastUserModel, err = service.repository.GetLastId(tx)
	helper.PanicIfError(err)

	defaultPassword, _ := bcrypt.GenerateFromPassword([]byte("welcome1"), 10)

	// Parse dto to models users
	userModel := models.User{
		ID:       lastUserModel.ID + 1,
		UserCode: helper.RandomString(5),
		IsActive: func() *bool {
			b := true
			return &b
		}(),
		Password: string(defaultPassword),
	}
	err = smapping.FillStruct(&userModel, smapping.MapFields(&userRegister))
	helper.PanicIfError(err)

	err = service.repository.Register(tx, userModel)
	helper.PanicIfError(err)

	return err
}
