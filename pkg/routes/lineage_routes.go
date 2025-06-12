package routes

import (
	"github.com/gin-gonic/gin"
	"tov_tools/pkg/api"
)

func RegisterLineageRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/lineages", api.GetAllLineages)
		v1.GET("/lineages/:name", api.GetLineageByName)
	}
}
