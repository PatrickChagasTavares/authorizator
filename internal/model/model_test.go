package model

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ModelSuite struct {
	suite.Suite
}

func TestModelSuite(t *testing.T) {
	suite.Run(t, new(ModelSuite))
}
