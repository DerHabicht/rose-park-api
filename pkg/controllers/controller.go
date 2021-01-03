package controllers

import (
	"github.com/sirupsen/logrus"

	"github.com/derhabicht/rose-park-api/pkg/data"
)

type Controller struct {
	DB data.IConnection
	Logger *logrus.Logger
}

func NewController(db data.IConnection, logger *logrus.Logger) Controller {
	return Controller{
		DB: db,
		Logger: logger,
	}
}

