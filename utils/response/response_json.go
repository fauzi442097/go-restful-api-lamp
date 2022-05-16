package ResponseJson

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   uint   `json:"code"`
	Status string `json:"status"`
}

type ResponseSuccess struct {
	Response
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Response
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"` // bisa jadi tidak ada detail
}

func Success(c *gin.Context, statusCode int, data interface{}) {

	res := ResponseSuccess{}
	res.Code = uint(statusCode)
	res.Status = "Success"
	res.Data = data

	c.JSON(statusCode, res)
}

func NotFound(c *gin.Context) {

	res := Response{}
	res.Code = http.StatusNotFound
	res.Status = http.StatusText(http.StatusNotFound)

	c.JSON(http.StatusNotFound, res)
}
