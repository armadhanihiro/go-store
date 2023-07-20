package handler

import (
	"gostore/internal/pkg/reason"
	"gostore/internal/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handle response error
func ResponseError(c *gin.Context, statusCode int, message string) {
	resp := ResponseBody{
		Status:  "error",
		Message: message,
	}
	c.JSON(statusCode, resp)
}

// Handle response success
func ResponseSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	resp := ResponseBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, resp)
}

// Parse request data & validate struct
func BindAndCheck(c *gin.Context, data interface{}) bool {
	err := c.Bind(data)
	if err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return true
	}

	isError := validator.Check(data)
	if isError {
		ResponseError(c, http.StatusUnprocessableEntity, reason.RequestFormError)
		return true
	}

	return false
}
