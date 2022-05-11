package repositories

import (
	"go-restful-api-lamp/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetById(customerId uint, tx *gorm.DB) (models.Customer, error)
}

type customerRepositoryImpl struct {
}

// Consctrutor Customer Repository
func NewCustomerRepository() CustomerRepository {
	return &customerRepositoryImpl{}
}

func (r *customerRepositoryImpl) GetById(customerId uint, tx *gorm.DB) (models.Customer, error) {

	customer := models.Customer{}
	// err := tx.First(&customer, customerId).Error
	err := tx.Where("id = ?", customerId).First(&customer).Error

	return customer, err
}
