package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/pkg/config"
)

func StaticRoutes(route *gin.Engine, config *config.Config) {
	route.Static("/public", config.FolderConfig.PublicPath)
}
