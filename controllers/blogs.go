package controllers

import (
	"fmt"
	"github.com/derhabicht/rose-park/database"
	"github.com/derhabicht/rose-park/models"
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type BlogsController struct{}

func NewBlogsController() BlogsController {
	return BlogsController{}
}

// Blog creation handler.
// @Summary Create a new blog site.
// @Description Creates a new blog site to be managed by this backend. Authentication is required to use this endpoint.
// @Success 201 {object} models.Blog
// @Failure 401 {object} controllers.ControllerError Returns if an invalid token is provided with the request.
// @Failure 422 {object} controllers.ControllerError Returns if a site with this URL is already being managed.
// @Failure 500 {object} controllers.ControllerError Returns if there is an unknown issue with the database.
// @Router /sites [post]
func (bc BlogsController) Create(c *gin.Context) {
	var blog models.Blog

	if err := c.ShouldBind(&blog); err != nil {
		logrus.WithFields(logrus.Fields{
			"endpoint": "/api/blogs/v1/sites/create",
			"error":    err,
		}).Error("failed to bind Blog object")
		c.AbortWithStatusJSON(http.StatusBadRequest, ControllerError{"error": "request body is malformed"})
		return
	}

	r := database.DB.Create(&blog)
	if r.Error != nil {
		logrus.WithFields(logrus.Fields{
			"endpoint": "POST /api/blogs/v1/sites/",
			"error":    r.Error,
		}).Error("failed to create Blog object")
		if ve, ok := r.Error.(validation.Errors); ok {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				ControllerError{"field_errors": ve},
			)
		} else if r.Error.Error() == "pq: duplicate key value violates unique constraint \"blogs_url_key\"" {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				ControllerError{
					"error": fmt.Sprintf("A blog at %s already exists", blog.Domain),
				},
			)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ControllerError{"error": "failed to create blog"})
		}
		return
	}

	c.JSON(http.StatusCreated, blog)
}

// Blog list handler.
// @Summary List all blog sites managed by this backend.
// @Description Lists all blogs that are managed by this backend. Authentication is required to use this endpoint.
// @Success 200 {object} []models.Blog
// @Success 204 {string} Returns if there are no blogs registered with this backend.
// @Failure 401 {object} controllers.ControllerError Returns if an invalid token is provided with the request.
// @Failure 500 {object} controllers.ControllerError Returns if there is an unknown issue with the database.
// @Router /sites [get]
func (bc BlogsController) List(c *gin.Context) {
	var blogs []models.Blog

	r := database.DB.Find(&blogs)

	if r.Error != nil {
		logrus.WithFields(logrus.Fields{
			"endpoint": "GET /api/blogs/v1/sites/",
			"error":    r.Error,
		}).Error("Failed to retrieve Blog objects from database")
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ControllerError{"error": "an error occurred while processing request"},
		)
		return
	}

	if r.RowsAffected == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, blogs)
	}
}

// Blog fetch handler.
// @Summary Fetch data regarding this blog.
// @Description Retrieve the blog, all authors that write for this blog and their bios, and the ten most recent post
// @Description titles (with their slugs) published on this blog. Authentication is not required for this endpoint.
// @Success 200 {object} models.Blog
// @Failure 401 {object} controllers.ControllerError Returns if an invalid token is provided with the request.
// @Failure 404 {object} controllers.ControllerError Returns if no blog exists at the given domain.
// @Failure 500 {object} controllers.ControllerError Returns if there is an unknown issue with the database.
// @Router /sites/{domain} [get]
func (bc BlogsController) Fetch(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Endpoint not yet implemented.")
}

func (bc BlogsController) Update(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Endpoint not yet implemented.")
}

func (bc BlogsController) Delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Endpoint not yet implemented.")
}
