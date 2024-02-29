package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (bc *BaseController) SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: "success",
		Data:    data,
	})
}

func (bc *BaseController) ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

func (bc *BaseController) MessageResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
	})
}
