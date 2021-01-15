package controllers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestHealthControllerSuite struct {
	suite.Suite
}

func (suite *TestHealthControllerSuite) SetupTest() {

}

func (suite *TestHealthControllerSuite) TestHealthCheck_Returns200() {
	suite.T().Skip()
}

func TestHealthController(t *testing.T) {
	suite.Run(t, new(TestHealthControllerSuite))
}
