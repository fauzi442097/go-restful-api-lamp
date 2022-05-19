package services

import (
	"go-restful-api-lamp/dto"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/models"
	"go-restful-api-lamp/repositories"
	"go-restful-api-lamp/utils/transaction"

	"github.com/go-playground/validator/v10"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type CustomerService interface {
	GetById(customerId uint) dto.CustomerResponse
	GetAll() []dto.CustomerResponse
	DeleteById(customerId uint)
	Create(customer dto.CustomerRequest)
}

type customerServiceImpl struct {
	repository repositories.CustomerRepository
	db         *gorm.DB
	Validate   *validator.Validate
}

// NewCustomerService returns new CustomerService.
func NewCustomerService(repo repositories.CustomerRepository, db *gorm.DB, validate *validator.Validate) CustomerService {
	return &customerServiceImpl{
		repository: repo,
		db:         db,
		Validate:   validate,
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

func (service *customerServiceImpl) GetAll() []dto.CustomerResponse {

	tx := service.db.Begin()
	defer transaction.CommitOrRollback(tx)

	customers, err := service.repository.GetAll(tx)
	helper.PanicIfError(err)

	customersResponse := []dto.CustomerResponse{}

	for _, customer := range customers {
		customerResponse := dto.CustomerResponse{}
		err = smapping.FillStruct(&customerResponse, smapping.MapFields(&customer))
		helper.PanicIfError(err)
		customersResponse = append(customersResponse, customerResponse)
	}

	return customersResponse
}

func (service *customerServiceImpl) DeleteById(customerId uint) {

	tx := service.db.Begin()
	defer transaction.CommitOrRollback(tx)

	_, err := service.repository.GetById(customerId, tx)
	helper.PanicIfError(err)

	service.repository.DeleteById(customerId, tx)

}

func (service *customerServiceImpl) Create(customer dto.CustomerRequest) {

	err := service.Validate.Struct(customer)
	helper.PanicIfError(err)

	tx := service.db.Begin()
	defer transaction.CommitOrRollback(tx)

	cust := customer

	customerModel := models.Customer{}
	err = smapping.FillStruct(&customerModel, smapping.MapFields(&cust))
	helper.PanicIfError(err)

	err = service.repository.Create(tx, customerModel)
	helper.PanicIfError(err)
}
