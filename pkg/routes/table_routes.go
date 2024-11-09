// pkg/routes/table_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"tov_tools/pkg/api"
)

func RegisterTableRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/table/get", api.GetTable)
	}
}
