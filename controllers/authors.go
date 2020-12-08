package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorsController struct{}

func NewAuthorsController() AuthorsController {
	return AuthorsController{}
}

func (ac AuthorsController) Create(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}

// TODO: Require blog query param
func (ac AuthorsController) List(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}

func (ac AuthorsController) Fetch(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}

func (ac AuthorsController) Update(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}

func (ac AuthorsController) Delete(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ControllerError{"error": "endpoint not implemented"})
}
