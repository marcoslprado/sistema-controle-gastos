package services

import (
	"errors"

	"controle-gastos-backend/models" // Import das structs da pasta models //
)

var pessoas []models.Pessoa // Variável global para armazenar pessoas //
var NextID = 1              // Variável que vai ditar qual o ID do próximo User cadastrado //

func CadastrarPessoa(nome string, idade int) models.Pessoa { // Funcao que cadastra pessoas //
	pessoa := models.Pessoa{ // Criação de uma pessoa com seus atributos //
		ID:    NextID,
		Nome:  nome,
		Idade: idade,
	}
	pessoas = append(pessoas, pessoa) // Adiciona a pessoa na variável de "memoria" para armazenar esses dados //
	NextID++                          // Incrementa o ID para que a proxima pessoa cadastrada tenha o ID diferente da atual //
	return pessoa
}

func DeletarPessoa(id int) error {
	for i, pessoa := range pessoas {
		if pessoa.ID == id {
			// Remove a pessoa//
			pessoas = append(pessoas[:i], pessoas[i+1:]...)
			// Remove as transacoes dessa pessoa //
			RemoverTransacoesPorPessoa(id)
			return nil
		}
	}
	return errors.New("Pessoa nao encontrada")
}

// Retorna todas as pessoas cadastradas //
func ListarPessoas() []models.Pessoa {
	return pessoas
}

// Retorna uma pessoa pelo ID
func BuscarPessoaPorID(id int) (models.Pessoa, error) {
	for _, pessoa := range pessoas {
		if pessoa.ID == id {
			return pessoa, nil
		}
	}
	return models.Pessoa{}, errors.New("pessoa não encontrada")
}
