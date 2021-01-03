package config

import (
	"fmt"
	"strings"
)

// Environment is an enum that describes one of the four environments in which the API could be operating.
type Environment int
const (
	DEVELOPMENT = iota
	TEST
	STAGE
	PRODUCTION
)

// ParseEnvironment is used to determine the Environment enum value from the given string.
func ParseEnvironment(s string) (Environment, error) {
	switch strings.ToLower(s) {
	case "development":
		return DEVELOPMENT, nil
	case "test":
		return TEST, nil
	case "stage":
		return STAGE, nil
	case "production":
		return PRODUCTION, nil
	default:
		return -1, fmt.Errorf("%s is not a valid environment", s)
	}
}

// String converts an Environment enum value into a string.
func (e Environment) String() string {
	return [...]string{"development", "test", "stage", "production"}[e]
}

