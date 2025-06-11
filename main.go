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
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(logger)
	zap.ReplaceGlobals(logger)

	router := gin.New()
	// gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Recovery(), middleware.ZapLogger())
	
	routes.RegisterStaticRoutes(router)

	routes.RegisterDiceRoutes(router)
	routes.RegisterCharacterRoutes(router)
	routes.RegisterTableRoutes(router)
	routes.RegisterHeritageRoutes(router)

	log.Println("Server started at :8080")
	log.Fatal(router.Run(":8080"))
}
