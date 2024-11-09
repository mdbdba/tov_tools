package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tov_tools/pkg/character"
)

func CreateCharacter(c *gin.Context) {
	var char character.Character
	if err := c.ShouldBindJSON(&char); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// character.Create(&char)
	c.JSON(http.StatusOK, char)
}
