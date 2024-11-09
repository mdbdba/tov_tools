// pkg/routes/dice_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"tov_tools/pkg/api"
)

func RegisterDiceRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/dice/roll", api.RollDice)
	}
}
