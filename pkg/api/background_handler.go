package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tov_tools/pkg/character"
)

// GetBackgroundByName handles requests to retrieve background information by name
func GetBackgroundByName(c *gin.Context) {
	name := c.Param("name")

	lineage, err := character.GetBackgroundByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lineage)
}

// GetAllBackgrounds handles requests to retrieve all available Backgrounds
func GetAllBackgrounds(c *gin.Context) {

	response := gin.H{
		"backgrounds": make([]string, 0, len(character.Backgrounds)),
	}

	for _, background := range character.Backgrounds {
		response["backgrounds"] = append(response["backgrounds"].([]string), background.Name)
	}

	c.JSON(http.StatusOK, response)
}
