package routes

import (
	"github.com/gin-gonic/gin"

	userConstructor "github.com/phn00dev/go-crud/internal/domain/user/constructor"
)

func UserRoutes(route *gin.Engine) {
	// auth user route
	userApiRoute := route.Group("/v1/api/user")
	{
		authUserRoute := userApiRoute.Group("/auth")
		authUserRoute.POST("/register", userConstructor.UserAuthHandler.Register)
		authUserRoute.POST("/login", userConstructor.UserAuthHandler.Login)
	}
}
