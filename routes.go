package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/derhabicht/rose-park/controllers"
	"github.com/derhabicht/rose-park/middleware"
)

func configureResource(g *gin.RouterGroup, c controllers.Controller, r string) {
	g.POST(fmt.Sprintf("/%s", r), c.Create)
	g.GET(fmt.Sprintf("/%s", r), c.List)
	g.GET(fmt.Sprintf("/%s/:public_id", r), c.Fetch)
	g.PUT(fmt.Sprintf("/%s/:public_id", r), c.Update)
	g.DELETE(fmt.Sprintf("/%s/:public_id", r), c.Delete)
}

func newRouter(version string, logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(ginlogrus.Logger(logger), gin.Recovery())

	validator := middleware.GetValidator()
	router.Use(middleware.Authorize(validator))

	// Visit {host}/api/v1/swagger/index.html to see the API documentation.
	v1 := router.Group("/api/v1")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// {host}/api/v1/health resource
	health := controllers.NewHealthController(version)
	{
		v1.GET("/health", health.Check)
	}

	return router
}
