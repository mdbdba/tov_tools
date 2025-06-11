package routes

import (
	"github.com/gin-gonic/gin"
	"tov_tools/pkg/api"
)

func RegisterHeritageRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/heritages", api.GetAllHeritages)
		v1.GET("/heritages/:name", api.GetHeritageByName)
		v1.GET("/heritages/lineages", api.GetHeritagesByLineage)
	}
}
