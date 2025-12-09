package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrCode string

type AppErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, status int, message string, data any) {
	ctx.JSON(status, ApiResponse{
		Code:    status,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, status int, err string, detail any) {
	ctx.JSON(status, AppErr{
		Code:    status,
		Message: err,
		Detail:  detail,
	})
}
