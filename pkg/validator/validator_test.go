package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type ValidatorSuite struct {
	suite.Suite
}

type TestData struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=18"`
}

func TestValidatorSuite(t *testing.T) {
	suite.Run(t, new(ValidatorSuite))
}

func (suite *ValidatorSuite) TestValidStructSuccess() {
	testData := TestData{Name: "John Doe", Age: 25}

	err := ValidStruct(testData)

	suite.NoError(err)
}

func (suite *ValidatorSuite) TestValidStructError() {
	testData := TestData{Name: "", Age: 15}

	err := ValidStruct(testData)

	suite.Error(err)

	validationErr, ok := err.(validator.ValidationErrors)
	suite.True(ok)
	suite.Len(validationErr, 2) // Espera-se que haja dois erros de validação
}
