package response

import "github.com/gin-gonic/gin"

type successResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context, status int, message string, data any) {
	ctx.JSON(status, &successResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
