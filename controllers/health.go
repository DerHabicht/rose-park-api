package controllers

import (
	"github.com/derhabicht/rose-park/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/derhabicht/rose-park/database"
)

type ServiceStatus struct {
	Services map[string]bool `json:"status"`
	Version  string          `json:"version"`
	Errors   []string        `json:"errors,omitempty"`
}

type HealthController struct {
	Status ServiceStatus
}

// NewHealthController initializes a HealthController.
func NewHealthController(version string) HealthController {
	return HealthController{
		Status: ServiceStatus{
			Services: make(map[string]bool),
			Version:  version,
		},
	}
}

// Healthcheck handler.
// @Summary Check to assure that the service is running.
// @Description Healthcheck endpoint. Reports which statuses are currently
// @Description running and the current API\'s version number. If critical
// @Description services are running, it will return 200. If any of the
// @Description critical services are down, then the endpoint will return 503.
// @Success 200 {object} controllers.ServiceStatus
// @Failure 503 {object} controllers.ServiceStatus
// @Router /health [get]
func (h HealthController) Check(c *gin.Context) {
	httpStatus := http.StatusOK

	h.Status.Services["endpoint"] = true

	err := database.DB.DB().Ping()
	if err == nil {
		h.Status.Services["database"] = true
	} else {
		h.Status.Services["database"] = false
		h.Status.Errors = append(h.Status.Errors, err.Error())
		httpStatus = http.StatusServiceUnavailable
	}

	config.Log.WithFields(logrus.Fields{
		"status": MarshalForLog(h.Status),
	}).Debug("Health check endpoint called.")

	c.JSON(httpStatus, h.Status)
}
