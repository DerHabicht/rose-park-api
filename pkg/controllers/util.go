package controllers

import "encoding/json"

// MarshalForLog is a utility function used to marshal structs to be included in log entries.
func MarshalForLog(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "FAILED TO MARSHAL OBJECT"
	}

	return string(b)
}
