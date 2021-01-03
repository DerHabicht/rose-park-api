package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/derhabicht/rose-park-api/pkg/controllers"
)

type ServiceStatus struct {
	Services map[string]bool `json:"status"`
	Version  string          `json:"version"`
	Errors   []string        `json:"errors,omitempty"`
}

type HealthController struct {
	controllers.Controller
	Status ServiceStatus
}

// NewHealthController initializes a HealthController.
func NewHealthController(controller controllers.Controller, version string) HealthController {
	return HealthController{
		Controller: controller,
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
func (t HealthController) Check(c *gin.Context) {
	httpStatus := http.StatusOK

	t.Status.Services["endpoint"] = true

	err := t.DB.Ping()
	if err == nil {
		t.Status.Services["database"] = true
	} else {
		t.Status.Services["database"] = false
		t.Status.Errors = append(t.Status.Errors, err.Error())
		httpStatus = http.StatusServiceUnavailable
	}

	t.Logger.WithFields(logrus.Fields{
		"status": controllers.MarshalForLog(t.Status),
	}).Debug("Health check endpoint called.")

	c.JSON(httpStatus, t.Status)
}
