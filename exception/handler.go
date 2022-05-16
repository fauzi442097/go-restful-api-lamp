package exception

import (
	"errors"
	"fmt"

	ResponseJson "go-restful-api-lamp/utils/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorHandler(c *gin.Context, recovered interface{}) {

	fmt.Println("Panic : ", recovered)
	if isNotFound := errors.Is(recovered.(error), gorm.ErrRecordNotFound); isNotFound {
		ResponseJson.NotFound(c)
	}

}
