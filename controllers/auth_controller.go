package controllers

import (
	"go-restful-api-lamp/dto"
	"go-restful-api-lamp/helper"
	"go-restful-api-lamp/services"
	ResponseJson "go-restful-api-lamp/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) AuthController {
	return &AuthControllerImpl{service}
}

func (controller *AuthControllerImpl) Login(c *gin.Context) {

	credential := dto.LoginRequest{}
	err := c.ShouldBindJSON(&credential)
	helper.PanicIfError(err)

	data := controller.service.Login(credential)

	ResponseJson.Success(c, http.StatusOK, nil, data)
}

func (controller *AuthControllerImpl) Register(c *gin.Context) {

	userRegister := dto.RegisterRequest{}
	err := c.ShouldBindJSON(&userRegister)
	helper.PanicIfError(err)

	err = controller.service.Register(userRegister)
	if err != nil {
		ResponseJson.Error(c, http.StatusInternalServerError, err.Error(), nil)
	}

	ResponseJson.Success(c, http.StatusOK, "Success Registered", nil)

}
