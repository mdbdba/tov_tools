// pkg/routes/character_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"tov_tools/pkg/api"
)

func RegisterCharacterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/character/create", api.CreateCharacter)
	}
}
