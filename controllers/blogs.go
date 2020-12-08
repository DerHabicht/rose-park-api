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
// @Failure 422 {object} controllers.ControllerError Returns if there are problems with the provided object to save.
// @Failure 500 {object} controllers.ControllerError Returns if there is an unknown issue with the database.
// @Router /sites [post]
func (bc BlogsController) Create(c *gin.Context) {
	var blog models.Blog

	if err := c.ShouldBind(&blog); err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "controllers",
			"endpoint": "POST /api/blogs/v1/sites/",
			"error":    err,
		}).Error("failed to bind Blog object")
		c.AbortWithStatusJSON(http.StatusBadRequest, ControllerError{"error": "request body is malformed"})
		return
	}

	r := database.DB.Create(&blog)
	if r.Error != nil {
		logrus.WithFields(logrus.Fields{
			"module": "controllers",
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
			"module": "controllers",
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

// Blog update handler.
// @Summary Edit data pertaining to a managed blog.
// @Description This endpoint is used to update the name or domain name of a managed blog. Authentication is required.
// @Success 200 {object} models.Blog
// @Failure 401 {object} controllers.ControllerError Returns if an invalid token is provided with the request.
// @Failure 404 {object} controllers.ControllerError Returns if the lookup domain doesn't point to a managed blog.
// @Failure 422 {object} controllers.ControllerError If there are problems with the updated object.
// @Failure 500 {object} controllers.ControllerError Returns if there is an unknown issue with the database.
// @Router /sites/{domain} [put]
func (bc BlogsController) Update(c *gin.Context) {
	var blog models.Blog

	r := database.DB.Where("domain = ?", c.Param("blog_domain")).First(&blog)
	if r.RowsAffected == 0 {
		logrus.WithFields(logrus.Fields{
			"module": "controllers",
			"endpoint": fmt.Sprintf("PUT /api/blogs/v1/sites/%s", c.Param("blog_domain")),
		}).Error("Blog to update was not found.")

		c.AbortWithStatusJSON(
			http.StatusNotFound,
			ControllerError{
				"error": fmt.Sprintf("%s is not managed by this backend.", c.Param("blog_domain")),
			},
		)

		return
	}

	logrus.WithFields(logrus.Fields{
		"module": "controllers",
		"endpoint": fmt.Sprintf("PUT /api/blogs/v1/sites/%s", c.Param("blog_domain")),
		"update_value": MarshalForLog(blog),
	}).Debug(fmt.Sprintf("Retrieved %s from database.", c.Param("blog_domain")))


	if err := c.ShouldBind(&blog); err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "controllers",
			"endpoint": fmt.Sprintf("PUT /api/blogs/v1/sites/%s", c.Param("blog_domain")),
			"error":    err,
		}).Error("failed to bind Blog object")
		c.AbortWithStatusJSON(http.StatusBadRequest, ControllerError{"error": "request body is malformed"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"module": "controllers",
		"endpoint": fmt.Sprintf("PUT /api/blogs/v1/sites/%s", c.Param("blog_domain")),
		"update_value": MarshalForLog(blog),
	}).Debug(fmt.Sprintf("%s will be updated.", c.Param("blog_domain")))

	r = database.DB.Save(&blog)
	if r.Error != nil {
		logrus.WithFields(logrus.Fields{
			"module": "controllers",
			"endpoint": fmt.Sprintf("PUT /api/blogs/v1/sites/%s", c.Param("blog_domain")),
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
			c.AbortWithStatusJSON(http.StatusInternalServerError, ControllerError{"error": "failed to update blog"})
		}
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (bc BlogsController) Delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Endpoint not yet implemented.")
}
