package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTable(c *gin.Context) {
	// Dummy implementation, replace with actual logic
	table := map[string]string{"example": "data"}
	c.JSON(http.StatusOK, table)
}
