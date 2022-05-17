package ResponseJson

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   uint   `json:"code"`
	Status string `json:"status"`
}

type ResponseSuccess struct {
	Response
	Data interface{} `json:"data,omitempty"`
}

type ResponseError struct {
	Response
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"` // bisa jadi tidak ada detail
}

var (
	defaultMessage       string
	defaultDetailMessage string
)

func Success(c *gin.Context, statusCode int, message interface{}, data interface{}) {

	if message != nil {
		defaultMessage = message.(string)
	} else {
		defaultMessage = "Success"
	}

	res := ResponseSuccess{}
	res.Code = uint(statusCode)
	res.Status = defaultMessage

	if data != nil {
		res.Data = data
	}

	c.JSON(statusCode, res)
}

func NotFound(c *gin.Context, message string) {

	res := ResponseError{}
	res.Code = http.StatusNotFound
	res.Status = http.StatusText(http.StatusNotFound)
	res.Message = message

	c.JSON(http.StatusNotFound, res)
}

func Error(c *gin.Context, statusCode int, message string, detail interface{}) {

	if message == "" {
		defaultMessage = http.StatusText(statusCode)
	} else {
		defaultMessage = message
	}

	fmt.Println(defaultMessage)

	res := ResponseError{}
	res.Code = uint(statusCode)
	res.Status = http.StatusText(statusCode)
	res.Message = defaultMessage

	if detail != "" {
		res.Detail = detail
	}

	c.JSON(statusCode, res)
}
