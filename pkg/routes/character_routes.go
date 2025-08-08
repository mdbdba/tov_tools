// pkg/routes/character_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"tov_tools/pkg/api"
)

func RegisterCharacterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Character creation
		v1.POST("/character/create", api.CreateCharacter)

		// Get character by name
		v1.GET("/character/name/:name", api.GetCharacterByName)

		// Get character by ID
		v1.GET("/character/id/:id", api.GetCharacterByID)

		// Update character by ID
		v1.PUT("/character/id/:id", api.UpdateCharacter)

		// Delete character by ID
		v1.DELETE("/character/id/:id", api.DeleteCharacter)

		// Get all characters
		v1.GET("/characters", api.GetAllCharacters)
	}
}
