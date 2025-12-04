package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrCode string

type AppErr struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`
}

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, ApiResponse{
		Code:    status,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, status int, err string, detail interface{}) {
	ctx.JSON(status, AppErr{
		Code:    status,
		Message: err,
		Detail:  detail,
	})
}
