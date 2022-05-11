package controllers

import (
	"fmt"
	"go-restful-api-lamp/services"
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
	if err != nil {
		panic(err)
	}

	customer := controller.service.GetById(uint(id))

	fmt.Println(customer)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Sukses",
		"data":    customer,
	})

}
