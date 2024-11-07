package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/utils/response"
	jwttoken "github.com/phn00dev/go-crud/pkg/jwt_token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(ctx, http.StatusUnauthorized, "error", "Authorization header required")
			ctx.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.Error(ctx, http.StatusUnauthorized, "error", "Bearer token required")
			ctx.Abort()
			return
		}
		claims, err := jwttoken.ValidateToken(tokenString)
		if err != nil {
			response.Error(ctx, http.StatusUnauthorized, "error", "Invalid token")
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claims.UserId)
		ctx.Next()
	}
}
