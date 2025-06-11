package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterStaticRoutes(router *gin.Engine) {
	// Handle favicon.ico requests to prevent 404 errors
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent) // Return 204 No Content
	})
}
