package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tov_tools/pkg/character"
)

// GetHeritageByName handles requests to retrieve heritage information by name
func GetHeritageByName(c *gin.Context) {
	name := c.Param("name")

	// Get heritage information
	heritage, err := character.GetHeritageByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, heritage)
}

// GetAllHeritages handles requests to retrieve all available heritages
func GetAllHeritages(c *gin.Context) {
	// Create a response with heritage names
	response := gin.H{
		"heritages": make([]string, 0, len(character.Heritages)),
	}

	// Extract all heritage names
	for _, heritage := range character.Heritages {
		response["heritages"] = append(response["heritages"].([]string), heritage.Name)
	}

	c.JSON(http.StatusOK, response)
}

// GetHeritagesByLineage handles requests to retrieve heritage suggestions by lineage
func GetHeritagesByLineage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"lineages": character.HeritageSuggestion(),
	})
}
