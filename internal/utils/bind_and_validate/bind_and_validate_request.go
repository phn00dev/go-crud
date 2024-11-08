package bindandvalidate

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/utils/response"
	"github.com/phn00dev/go-crud/internal/utils/validate"
)

func BindAndValidateRequest(ctx *gin.Context, request interface{}) bool {
	if err := ctx.ShouldBind(&request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "body parser error", err.Error())
		return false
	}
	if err := validate.ValidateStruct(request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "validate error", err.Error())
		return false
	}
	return true
}
