package main

import (
	"github.com/derhabicht/rose-park/middleware"

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

	validator := middleware.Authorize(middleware.GetValidator())

	// Visit {host}/api/v1/swagger/index.html to see the API documentation.
	v1 := router.Group("/api/blogs/v1")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// {host}/api/v1/health resource
	health := controllers.NewHealthController(version)
	{
		v1.GET("/health", health.Check)
	}

	// Blog setup/management endpoints
	blogs := controllers.NewBlogsController()
	v1.POST("/sites", validator, blogs.Create)
	v1.GET("/sites", validator, blogs.List)
	v1.PUT("/sites/:blog_domain", validator, blogs.Update)
	v1.DELETE("/sites/:blog_domain", validator, blogs.Delete)

	// Author bios endpoints
	authors := controllers.NewAuthorsController()
	v1.POST("/authors", validator, authors.Create)
	v1.GET("/authors", authors.List)
	v1.GET("/authors/:author_id", authors.Fetch)
	v1.PUT("/authors/:author_id", validator, authors.Update)
	v1.DELETE("/authors/:author_id", validator, authors.Delete)

	// Posts endpoints
	posts := controllers.NewPostsController()
	v1.POST("/posts/:blog_domain", validator, posts.Create)
	v1.GET("/posts/:blog_domain", posts.List)
	v1.GET("/posts/:blog_domain/:post_slug", posts.Fetch)
	v1.PUT("/posts/:blog_domain/:post_slug", validator, posts.Update)
	v1.DELETE("/posts/:blog_domain/:post_slg", validator, posts.Delete)

	return router
}
