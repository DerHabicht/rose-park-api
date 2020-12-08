package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostsController struct{}

func NewPostsController() PostsController {
	return PostsController{}
}

func (pc PostsController) Create(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}

func (pc PostsController) List(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}

func (pc PostsController) Fetch(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}

func (pc PostsController) Update(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}

func (pc PostsController) Delete(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}
