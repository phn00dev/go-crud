package routes

import (
	"github.com/gin-gonic/gin"

	postConstructor "github.com/phn00dev/go-crud/internal/domain/post/constructor"
	userConstructor "github.com/phn00dev/go-crud/internal/domain/user/constructor"
	"github.com/phn00dev/go-crud/internal/middleware"
)

func UserRoutes(route *gin.Engine) {
	// auth user route
	userApiRoute := route.Group("/v1/api/user")
	{
		authUserRoute := userApiRoute.Group("/auth")
		authUserRoute.POST("/register", userConstructor.UserAuthHandler.Register)
		authUserRoute.POST("/login", userConstructor.UserAuthHandler.Login)
		// use middleware
		userRoute := userApiRoute.Group("/user")
		{
			// user profile data
			userRoute.Use(middleware.AuthMiddleware())
			userRoute.GET("/", userConstructor.UserHandler.GetUser)
			userRoute.PUT("/update-data", userConstructor.UserHandler.Update)
			userRoute.PUT("/update-password", userConstructor.UserHandler.UpdatePassword)
			userRoute.DELETE("/delete", userConstructor.UserHandler.Delete)
			// user posts
			userPostRoute := userRoute.Group("/posts")
			{
				userPostRoute.GET("/", postConstructor.PostHandler.GetAll)
				userPostRoute.GET("/:postId", postConstructor.PostHandler.GetOne)
				userPostRoute.POST("/create", postConstructor.PostHandler.Create)
				userPostRoute.PUT("/:postId", postConstructor.PostHandler.Update)
				userPostRoute.DELETE("/:postId", postConstructor.PostHandler.Delete)
			}

		}

		publicPostRoute := userApiRoute.Group("/posts")
		{
			publicPostRoute.GET("/", postConstructor.PostHandler.GetAllPost)
			publicPostRoute.GET("/:post_slug", postConstructor.PostHandler.GetPostBySlug)
		}
	}

}
