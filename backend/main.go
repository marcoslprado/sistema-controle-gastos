package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"controle-gastos-backend/models"
	"controle-gastos-backend/services"

	"github.com/rs/cors"
)

func main() {
	// Configura o CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Permite o frontend
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Cria o roteador
	router := http.NewServeMux()

	// Endpoint para cadastrar uma pessoa
	router.HandleFunc("/pessoas", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var pessoa models.Pessoa
			err := json.NewDecoder(r.Body).Decode(&pessoa)
			if err != nil {
				http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
				return
			}
			if pessoa.Nome == "" || pessoa.Idade <= 0 {
				http.Error(w, "Nome e idade são obrigatórios", http.StatusBadRequest)
				return
			}
			pessoaCadastrada := services.CadastrarPessoa(pessoa.Nome, pessoa.Idade)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pessoaCadastrada)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint para listar todas as pessoas
	router.HandleFunc("/pessoas/listar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pessoas := services.ListarPessoas()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pessoas)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint para cadastrar uma transação
	router.HandleFunc("/transacoes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var transacao models.Transacao
			err := json.NewDecoder(r.Body).Decode(&transacao)
			if err != nil {
				http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
				return
			}
			err = services.CadastrarTransacao(transacao.Descricao, transacao.Valor, transacao.Tipo, transacao.Pessoa)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint para listar todas as transações
	router.HandleFunc("/transacoes/listar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			transacoes := services.ListarTransacoes()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(transacoes)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint para consultar totais
	router.HandleFunc("/totais", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			totais := services.CalcularTotaisPorPessoa()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(totais)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
	// Inicia o servidor com CORS
	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", c.Handler(router))
}
