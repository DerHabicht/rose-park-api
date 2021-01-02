package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/derhabicht/rose-park/controllers"
)

func newRouter(version string, logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(ginlogrus.Logger(logger), gin.Recovery())

	// Visit {host}/api/v1/swagger/index.html to see the API documentation.
	v1 := router.Group("/api/blogs/v1")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// {host}/api/v1/health resource
	health := controllers.NewHealthController(version)
	{
		v1.GET("/health", health.Check)
	}

	return router
}
