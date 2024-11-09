package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"tov_tools/pkg/middleware"
	"tov_tools/pkg/routes"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	router := gin.New()
	router.Use(gin.Recovery(), middleware.ZapLogger())

	routes.RegisterDiceRoutes(router)
	routes.RegisterCharacterRoutes(router)
	routes.RegisterTableRoutes(router)

	log.Println("Server started at :8080")
	log.Fatal(router.Run(":8080"))
}
