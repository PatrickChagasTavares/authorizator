# Autorizador de Negócios 

Esse projeto implementa um Autorizador de Negócios, que avalia a elegibilidade de um negócio com base em critérios específicos. O autorizador considera dados pessoais, pontuação de crédito e informações sobre o imóvel de interesse.

## 🚀 Começando

Siga estas instruções para configurar o projeto na sua máquina local para desenvolvimento e teste.


## Exemplo de uso:
<img width="1092" alt="exemple_authorizator" src="https://github.com/PatrickChagastavares/authorizator/assets/49497853/f7124d27-83b9-4628-8fc3-6978a5cb9b05">


### 📋 Pré-requisitos

Ferramentas necessárias:

- [Golang](https://golang.org/doc/install)

## 📦 Desenvolvimento

Comandos importantes para rodar o projeto e validar:

- `make run`: Compila e executa o código principal.
- `make test`: Executa os testes do projeto e mostra a cobertura.
- `make test-cover`: O mesmo do `make test`, porém abre o brawser para mais detalhes.
- `make help`: imprime os comando disponiveis no `make`

## 🗂 Estrutura do Projeto

### Descrição dos Pacotes e Arquivos Principais:

- `./cmd/autorizador/main.go`: O código que inicia a aplicação.
- `./pkg/reader`: Pacote para leitura de dados JSON.
- `./pkg/validator`: Pacote para validação de campos usando o validador.
- `./internal/authorization`: Lógica relacionada à autorização de negócios.
- `./internal/model`: Definições de structs relacionadas aos dados de entrada e saída.
- `./makefile`: Arquivo de make para automatizar tarefas comuns.

## 🛠️ Construído com

- [Golang](https://golang.org) - Linguagem de Programação
- [Validator](https://github.com/go-playground/validator/v10) - Validador de structs

## 📝 Licença

Este projeto é licenciado sob a licença MIT - consulte o arquivo [LICENSE](LICENSE) para obter detalhes.
