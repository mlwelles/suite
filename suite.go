package suite

import (
	"context"
	"math/rand"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestingSuite = suite.TestingSuite

// have your test suite extend this or StandaloneSuite
type Suite struct {
	suite.Suite
	Ctx           context.Context
	didSetupSuite bool
}

func (s *Suite) ensureSetup() {
	if !s.didSetupSuite {
		s.SetupSuite()
		s.didSetupSuite = true
	}
	return
}

func (s *Suite) SetupSuite() {
	rand.Seed(time.Now().UnixNano())
	s.Ctx = context.Background()
	s.didSetupSuite = true
}

func (s *Suite) SetupTest() {
	// Override in your test suite if needed
}

func (s *Suite) Log(args ...interface{}) {
	s.T().Log(args...)
}