package repositories

import (
	"go-restful-api-lamp/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetAll(tx *gorm.DB) ([]models.Customer, error)
	GetById(customerId uint, tx *gorm.DB) (models.Customer, error)
	DeleteById(customerId uint, tx *gorm.DB) error
	Create(tx *gorm.DB, customer models.Customer) error
}

type customerRepositoryImpl struct {
}

// Consctrutor Customer Repository
func NewCustomerRepository() CustomerRepository {
	return &customerRepositoryImpl{}
}

func (r *customerRepositoryImpl) GetById(customerId uint, tx *gorm.DB) (models.Customer, error) {

	customer := models.Customer{}
	err := tx.Where("id = ?", customerId).First(&customer).Error
	return customer, err

}

func (r *customerRepositoryImpl) GetAll(tx *gorm.DB) ([]models.Customer, error) {

	customers := []models.Customer{}

	// Get All Record
	err := tx.Limit(20).Order("id desc").Find(&customers).Error
	return customers, err

}

func (r *customerRepositoryImpl) DeleteById(customerId uint, tx *gorm.DB) error {

	customer := models.Customer{}
	err := tx.Where("id = ?", customerId).Delete(&customer).Error
	return err

}

func (r *customerRepositoryImpl) Create(tx *gorm.DB, customer models.Customer) error {

	err := tx.Create(&customer).Error
	return err

}
