package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginLogrus "github.com/toorop/gin-logrus"

	// _ "github.com/derhabicht/rose-park-api/api"
	instance "github.com/derhabicht/rose-park-api/internal/controllers"
	"github.com/derhabicht/rose-park-api/pkg/controllers"
	"github.com/derhabicht/rose-park-api/pkg/data"
)

func SetUpRouter(db data.IConnection, logger *logrus.Logger, version string) *gin.Engine {
	router := gin.New()
	router.Use(ginLogrus.Logger(logger), gin.Recovery())

	v1 := router.Group("/api/blogs/v1")

	// Swagger docs at /api/blogs/v1/swagger/index.html
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rootController := controllers.NewController(db, logger)

	healthController := instance.NewHealthController(rootController, version)
	health := v1.Group("/health")
	{
		health.GET("/", healthController.Check)
	}

	return router
}
