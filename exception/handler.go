package exception

import (
	"errors"
	"net/http"
	"reflect"

	ResponseJson "go-restful-api-lamp/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorHandler(c *gin.Context, recovered interface{}) {

	var err string
	validationErrors, isValidationError := recovered.(validator.ValidationErrors)
	if isValidationError {

		detailErrors := []map[string]string{}
		for _, validation := range validationErrors {

			detailError := map[string]string{
				"key":   validation.Field(),
				"error": validation.Error(),
			}

			detailErrors = append(detailErrors, detailError)
		}

		ResponseJson.Error(c, http.StatusBadRequest, "Invalid Request", detailErrors)
		return
	}

	if reflect.TypeOf(recovered).Name() == "string" {
		err = recovered.(string)
	} else {
		err = recovered.(error).Error()

		if isNotFound := errors.Is(recovered.(error), gorm.ErrRecordNotFound); isNotFound {
			ResponseJson.NotFound(c, err)
			return
		}

	}

	ResponseJson.Error(c, http.StatusInternalServerError, err, "")

}
