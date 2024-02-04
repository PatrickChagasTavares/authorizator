package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/patrickchagastavares/authorizator/internal/authorization"
	"github.com/patrickchagastavares/authorizator/internal/model"
	"github.com/patrickchagastavares/authorizator/pkg/reader"
	"github.com/patrickchagastavares/authorizator/pkg/validator"
)

type authorizator struct {
	Person   model.Person   `json:"pessoa" validate:"required"`
	Score    model.Score    `json:"score" validate:"required"`
	Property model.Property `json:"imovel" validate:"required"`
}

func main() {
	fmt.Println("Insira o JSON para análise:")

	var body authorizator
	if err := reader.ReadJSON(os.Stdin, &body); err != nil {
		panic("Erro falha ao ler JSON")
	}

	if err := validator.ValidStruct(body); err != nil {
		fmt.Printf("input.validator.error: %s", err)
		return
	}

	output := authorization.AutorizarNegocio(body.Person, body.Score, body.Property)

	// convert output to json and print
	b, err := json.Marshal(output)
	if err != nil {
		fmt.Println("Erro ao formatar saida: ", err)
		return
	}

	fmt.Printf("\n%s\n", string(b))
}
