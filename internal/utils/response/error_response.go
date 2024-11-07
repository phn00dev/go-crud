package response

import (
	"github.com/gin-gonic/gin"

)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

func Error(ctx *gin.Context, status int, message, errors string) {
	ctx.JSON(status, errorResponse{
		Status:  status,
		Message: message,
		Errors:  errors,
	})
}
