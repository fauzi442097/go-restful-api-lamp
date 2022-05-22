package middleware

import (
	"net/http"
	"strings"

	"go-restful-api-lamp/models"
	"go-restful-api-lamp/services"
	ResponseJson "go-restful-api-lamp/utils/response"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		ResponseJson.Error(c, http.StatusBadRequest, "Authorization token not found", nil)
		c.Abort()
		return
	}

	// Validate Token
	mapAuth := strings.Split(authorization, " ")
	token := mapAuth[1]

	claims := models.Claims{}
	_, errToken := services.ValidateToken(token, &claims)

	if errToken != nil {
		ResponseJson.Error(c, http.StatusUnauthorized, errToken.Error(), nil)
		c.Abort()
	}

	c.Next()
}
