package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tov_tools/pkg/character"
)

// GetLineageByName handles requests to retrieve lineage information by name
func GetLineageByName(c *gin.Context) {
	name := c.Param("name")

	// Get lineage information
	lineage, err := character.GetLineageByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lineage)
}

// GetAllLineages handles requests to retrieve all available lineages
func GetAllLineages(c *gin.Context) {
	// Create a response with lineage names
	response := gin.H{
		"lineages": make([]string, 0, len(character.Lineages)),
	}

	// Extract all lineage names
	for _, lineage := range character.Lineages {
		response["lineages"] = append(response["lineages"].([]string), lineage.Name)
	}

	c.JSON(http.StatusOK, response)
}
