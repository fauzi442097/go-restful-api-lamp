package controllers

import (
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/services"
	ResponseJson "go-restful-api-lamp/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	GetById(c *gin.Context)
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
	helper.PanicIfError(err)

	customer := controller.service.GetById(uint(id))

	ResponseJson.Success(c, http.StatusOK, customer)
}
