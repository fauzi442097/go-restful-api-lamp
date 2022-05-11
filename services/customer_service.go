package services

import (
	"fmt"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/models"
	"go-restful-api-lamp/repositories"

	"gorm.io/gorm"
)

type CustomerService interface {
	GetById(customerId uint) models.Customer
}

type customerServiceImpl struct {
	repository repositories.CustomerRepository
	db         *gorm.DB
}

// NewCustomerService returns new CustomerService.
func NewCustomerService(repo repositories.CustomerRepository, db *gorm.DB) CustomerService {
	return &customerServiceImpl{
		repository: repo,
		db:         db,
	}
}

func (service *customerServiceImpl) GetById(customerId uint) models.Customer {

	// Start Transaction
	tx := service.db.Begin()

	customer, err := service.repository.GetById(customerId, tx)
	defer helper.CommitOrRollback(tx)

	fmt.Println("error : ", err)
	if err != nil {
		err := tx.Rollback().RowsAffected
		fmt.Println("affected : ", err)
	} else {
		// Commit Transaction
		da := tx.Commit().Error
		fmt.Println("da", da)
	}

	return customer
	
}
