package services

import (
	"controle-gastos-backend/models"
	"errors"
)

var transacoes []models.Transacao // Variavel global para armazenar transacoes //
var nextTransacaoID = 1           // Variável que vai ditar qual o ID da proxima transacao cadastrada //

// Funcao que adiciona uma nova transacao//
func CadastrarTransacao(descricao string, valor float64, tipo models.Tipo, pessoaID int) error {
	// Verifica se a pessoa existe //
	pessoa, err := BuscarPessoaPorID(pessoaID)
	if err != nil {
		return errors.New("Pessoa não encontrada") // Retorna um erro no caso de a pessoa nao ser encontrada //
	}

	if pessoa.Idade < 18 && tipo != models.Despesa {
		return errors.New("Menor de idade só pode ter despesas") // Nao permite que um menor de idade cadastre receitas //
	}

	// Criação do objeto estruturado na struct Transacao //
	transacao := models.Transacao{
		ID:        nextTransacaoID,
		Descricao: descricao,
		Valor:     valor,
		Tipo:      tipo,
		Pessoa:    pessoaID,
	}
	transacoes = append(transacoes, transacao) // Adiciona a transacao no array de memoria //
	nextTransacaoID++                          // Incrementa o ID para que a proxima transacao cadastrada nao tenha o mesmo ID da atual //
	return nil
}

// Retorna todas as transações cadastradas //
func ListarTransacoes() []models.Transacao {
	return transacoes
}

// Remove todas as transações associadas a uma pessoa //
func RemoverTransacoesPorPessoa(pessoaID int) {
	var novasTransacoes []models.Transacao
	for _, transacao := range transacoes {
		if transacao.Pessoa != pessoaID { // Basicamente migra as transacoes que nao vao ser excluidas para um novo array //
			novasTransacoes = append(novasTransacoes, transacao)
		}
	}
	transacoes = novasTransacoes // Copia o novo array para o primeiro //
}

// Calcula o total de receitas, despesas e o saldo líquido por pessoa
func CalcularTotaisPorPessoa() map[int]map[string]float64 {
	totaisPorPessoa := make(map[int]map[string]float64)

	for _, transacao := range transacoes {
		pessoaID := transacao.Pessoa

		// Inicializa os totais para a pessoa, se necessário
		if _, ok := totaisPorPessoa[pessoaID]; !ok {
			totaisPorPessoa[pessoaID] = map[string]float64{
				"totalReceitas": 0,
				"totalDespesas": 0,
				"saldoLiquido":  0,
			}
		}

		// Atualiza os totais
		if transacao.Tipo == models.Receita {
			totaisPorPessoa[pessoaID]["totalReceitas"] += transacao.Valor
		} else if transacao.Tipo == models.Despesa {
			totaisPorPessoa[pessoaID]["totalDespesas"] += transacao.Valor
		}

		totaisPorPessoa[pessoaID]["saldoLiquido"] = totaisPorPessoa[pessoaID]["totalReceitas"] - totaisPorPessoa[pessoaID]["totalDespesas"]
	}

	return totaisPorPessoa
}
