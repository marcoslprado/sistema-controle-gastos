package models

// Tipo representa o tipo de transição: despesa ou receita //

type Tipo string

const (
	Despesa Tipo = "despesa"
	Receita      = "receita"
)

// Struct que representa uma Transação //

type Transacao struct {
	ID        int
	Descricao string
	Valor     float64
	Tipo      Tipo
	Pessoa    int
}
