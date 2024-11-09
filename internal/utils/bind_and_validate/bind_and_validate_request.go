package bindandvalidate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/phn00dev/go-crud/internal/utils/response"
	"github.com/phn00dev/go-crud/internal/utils/validate"
)

func BindAndValidateRequest(ctx *gin.Context, request interface{}) bool {
	// JSON, form-data, ýa-da multiparty binding üçin ShouldBind ulanylýar
	if err := ctx.ShouldBind(request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "body parser error", err.Error())
		return false
	}
	// Struct görnüşinde barlamak
	if err := validate.ValidateStruct(request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "validate error", err.Error())
		return false
	}
	return true
}

func BindAndValidateRequestofFormData(ctx *gin.Context, request any) bool {
	if err := ctx.ShouldBindWith(request, binding.Form); err != nil {
		response.Error(ctx, http.StatusBadRequest, "body parser error", err.Error())
		return false
	}
	if err := validate.ValidateStruct(request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "validate error", err.Error())
		return false
	}
	return true
}
