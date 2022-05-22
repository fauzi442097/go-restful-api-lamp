package routes

import (
	"go-restful-api-lamp/controllers"
	"go-restful-api-lamp/exception"
	"go-restful-api-lamp/middleware"
	"go-restful-api-lamp/repositories"
	"go-restful-api-lamp/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// var (
// 	customerRepository = repositories.NewcustomerRepository(db)
// 	customerService    = services.NewCustomerService(customerRepository)
// 	customerController = controllers.NewCustomerController(customerService)
// )

func Setup(db *gorm.DB, validator *validator.Validate) *gin.Engine {

	customerRepository := repositories.NewCustomerRepository()
	customerService := services.NewCustomerService(customerRepository, db, validator)
	customerController := controllers.NewCustomerController(customerService)

	authRepository := repositories.NewAuthRepository()
	authService := services.NewAuthService(authRepository, db, validator)
	authController := controllers.NewAuthController(authService)

	route := gin.New()

	route.Use(gin.Logger())
	route.Use(gin.CustomRecovery(exception.ErrorHandler))
	route.POST("/users/login", authController.Login)
	route.POST("/users/register", authController.Register)

	authorized := route.Use(middleware.JwtMiddleware)
	authorized.GET("/customers", customerController.GetAll)
	authorized.POST("/customers", customerController.Create)
	authorized.GET("/customers/:customerId", customerController.GetById)
	authorized.DELETE("/customers/:customerId", customerController.DeleteById)

	return route
}
