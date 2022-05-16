package services

import (
	"go-restful-api-lamp/dto"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/repositories"
	"go-restful-api-lamp/utils/transaction"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type CustomerService interface {
	GetById(customerId uint) dto.CustomerResponse
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

func (service *customerServiceImpl) GetById(customerId uint) dto.CustomerResponse {

	// Start Transaction
	tx := service.db.Begin()
	defer transaction.CommitOrRollback(tx)

	customerModel, err := service.repository.GetById(customerId, tx)
	helper.PanicIfError(err)

	customerResponse := dto.CustomerResponse{}
	err = smapping.FillStruct(&customerResponse, smapping.MapFields(&customerModel))
	helper.PanicIfError(err)

	return customerResponse
}
