package controllers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

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

// MarshalForLog is a utility function used to marshal structs to be included in log entries.
func MarshalForLog(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "FAILED TO MARSHAL OBJECT"
	}

	return string(b)
}