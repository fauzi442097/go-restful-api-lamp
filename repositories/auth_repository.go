package repositories

import (
	"go-restful-api-lamp/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(tx *gorm.DB, user models.User) (models.User, error)
	Register(tx *gorm.DB, user models.User) error
}

type AuthRepositoryImpl struct {
}

// NewAuthRepository returns AuthRepository(Interface).
func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (r *AuthRepositoryImpl) Login(tx *gorm.DB, user models.User) (models.User, error) {

	userModel := models.User{}
	err := tx.Where("username = ?", user.Username).First(&userModel).Error
	return userModel, err

}

func (r *AuthRepositoryImpl) Register(tx *gorm.DB, user models.User) error {

	err := tx.Create(&user).Error
	return err
}
