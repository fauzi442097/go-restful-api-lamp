package routes

import (
	"go-restful-api-lamp/controllers"
	"go-restful-api-lamp/repositories"
	"go-restful-api-lamp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var (
// 	customerRepository = repositories.NewcustomerRepository(db)
// 	customerService    = services.NewCustomerService(customerRepository)
// 	customerController = controllers.NewCustomerController(customerService)
// )

func Setup(db *gorm.DB) *gin.Engine {

	customerRepository := repositories.NewCustomerRepository()
	customerService := services.NewCustomerService(customerRepository, db)
	customerController := controllers.NewCustomerController(customerService)

	route := gin.Default()
	route.GET("/customers/:customerId", customerController.GetById)

	return route
}
