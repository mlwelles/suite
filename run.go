package suite

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func Run(t *testing.T, s suite.TestingSuite) {
	suite.Run(t, s)
}
