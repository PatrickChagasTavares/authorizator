package authorization

import (
	"github.com/patrickchagastavares/authorizator/internal/model"
)

type Output struct {
	CPF         string   `json:"cpf"`
	Status      string   `json:"status"`
	Constraints []string `json:"constraints"`
}

func AutorizarNegocio(person model.Person, score model.Score, property model.Property) (response Output) {
	response = Output{CPF: person.Document}

	if !score.ValidScore() {
		response.Status = "reprovado"
		response.Constraints = append(response.Constraints, "score-muito-baixo")
	}

	if !person.ValidIncome() {
		response.Status = "reprovado"
		response.Constraints = append(response.Constraints, "renda-insuficiente")
	}

	if len(response.Constraints) > 0 {
		return
	}

	if !property.ValidMinInput() {
		response.Status = "reprovado-parcialmente"
		response.Constraints = append(response.Constraints, "entrada-insuficiente")
	}

	if !property.ValidHighInput() {
		response.Status = "aprovado"
		response.Constraints = append(response.Constraints, "entrada-elevada")
	}

	if len(response.Constraints) == 0 {
		response.Status = "aprovado"
		response.Constraints = append(response.Constraints, "-")
	}

	return
}
