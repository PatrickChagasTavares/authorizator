package authorization

import (
	"testing"

	"github.com/patrickchagastavares/authorizator/internal/model"
	"github.com/stretchr/testify/suite"
)

type AuthorizationSuite struct {
	suite.Suite
}

func TestAuthorizationSuite(t *testing.T) {
	suite.Run(t, new(AuthorizationSuite))
}

func (suite *AuthorizationSuite) TestAutorizarNegocioAproved() {
	person := model.Person{Name: "John Doe", Document: "12345678901", Income: 2500}
	score := model.Score{Punctuation: 600, Source: "serasa"}
	property := model.Property{Value: 100000, Input: 7000}

	output := AutorizarNegocio(person, score, property)

	suite.Equal("aprovado", output.Status)
	suite.Contains(output.Constraints, "-")
}

func (suite *AuthorizationSuite) TestAutorizarNegocioPartialRejection() {
	person := model.Person{Name: "Jane Doe", Document: "98765432101", Income: 2000}
	score := model.Score{Punctuation: 700, Source: "serasa"}
	property := model.Property{Value: 80000, Input: 3000}

	output := AutorizarNegocio(person, score, property)

	suite.Equal("reprovado-parcialmente", output.Status)
	suite.Contains(output.Constraints, "entrada-insuficiente")
}

func (suite *AuthorizationSuite) TestAutorizarNegocioHighInputApproval() {
	person := model.Person{Name: "Alice Johnson", Document: "56789012345", Income: 5000}
	score := model.Score{Punctuation: 800, Source: "serasa"}
	property := model.Property{Value: 200000, Input: 40000}

	output := AutorizarNegocio(person, score, property)

	suite.Equal("aprovado", output.Status)
	suite.Contains(output.Constraints, "entrada-elevada")
}

func (suite *AuthorizationSuite) TestAutorizarNegocioRejection() {
	person := model.Person{Name: "Bob Smith", Document: "45678901234", Income: 1500}
	score := model.Score{Punctuation: 400, Source: "serasa"}
	property := model.Property{Value: 120000, Input: 10000}

	output := AutorizarNegocio(person, score, property)

	suite.Equal("reprovado", output.Status)
	suite.Contains(output.Constraints, "score-muito-baixo")
	suite.Contains(output.Constraints, "renda-insuficiente")
}
