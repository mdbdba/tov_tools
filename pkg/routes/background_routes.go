package routes

import (
	"github.com/gin-gonic/gin"
	"tov_tools/pkg/api"
)

func RegisterBackgroundRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/backgrounds", api.GetAllBackgrounds)
		v1.GET("/backgrounds/:name", api.GetBackgroundByName)
	}
}
