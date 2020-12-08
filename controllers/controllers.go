package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Controller interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Fetch(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ControllerError map[string]interface{}

func (ce ControllerError) Error() string {
	s, err := json.Marshal(ce)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"map": ce,
		}).Error("Failed to marshal a ControllerError.")
	}

	return string(s)
}
