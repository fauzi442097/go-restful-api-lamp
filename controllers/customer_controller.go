package controllers

import (
	"go-restful-api-lamp/dto"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/services"
	ResponseJson "go-restful-api-lamp/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	DeleteById(c *gin.Context)
	Create(c *gin.Context)
}

type customerControllerImpl struct {
	service services.CustomerService
}

// NewCustomerController returns new CustomerController.
func NewCustomerController(s services.CustomerService) CustomerController {
	return &customerControllerImpl{s}
}

func (controller *customerControllerImpl) GetById(c *gin.Context) {

	customerId := c.Param("customerId")
	id, err := strconv.ParseUint(customerId, 10, 64)
	if err != nil {
		ResponseJson.Error(c, http.StatusBadRequest, "Customer id must be a number", err.(error).Error())
		return
	}

	customer := controller.service.GetById(uint(id))
	ResponseJson.Success(c, http.StatusOK, nil, customer)
}

func (controller *customerControllerImpl) GetAll(c *gin.Context) {

	customer := controller.service.GetAll()
	ResponseJson.Success(c, http.StatusOK, nil, customer)

}

func (controller *customerControllerImpl) DeleteById(c *gin.Context) {

	customerId := c.Param("customerId")
	id, err := strconv.ParseUint(customerId, 10, 64)
	if err != nil {
		ResponseJson.Error(c, http.StatusBadRequest, "Customer id must be a number", err.Error())
		return
	}

	controller.service.DeleteById(uint(id))
	ResponseJson.Success(c, http.StatusOK, "Customer successfully deleted", nil)
}

func (controller *customerControllerImpl) Create(c *gin.Context) {

	customerRequest := dto.CustomerRequest{}
	err := c.ShouldBindJSON(&customerRequest)
	helper.PanicIfError(err)

	controller.service.Create(customerRequest)

	ResponseJson.Success(c, http.StatusOK, nil, nil)

}
