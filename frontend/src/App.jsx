import React, { useState, useEffect } from "react";
import FormPessoa from "./components/FormPessoa";
import FormTransacao from "./components/FormTransacao";
import Totais from "./components/Totais";

function App() {
  const [pessoas, setPessoas] = useState([]);
  const [transacoes, setTransacoes] = useState([]);
  const [totais, setTotais] = useState({});
  const [loading, setLoading] = useState(true); // Estado para controlar o carregamento

  // Carrega pessoas e transações ao iniciar
  useEffect(() => {
    fetch("http://localhost:8080/pessoas/listar")
      .then((response) => response.json())
      .then((data) => {
        setPessoas(data || []); // Garante que pessoas seja um array, mesmo se data for null
      });

    fetch("http://localhost:8080/transacoes/listar")
      .then((response) => response.json())
      .then((data) => {
        setTransacoes(data || []); // Garante que transacoes seja um array, mesmo se data for null
        calcularTotaisPorPessoa(data || []); // Calcula os totais por pessoa ao carregar as transações
      });

    // Carrega os totais ao iniciar
    fetch("http://localhost:8080/totais")
      .then((response) => response.json())
      .then((data) => {
        setTotais(data || {}); // Garante que totais seja um objeto, mesmo se data for null
        setLoading(false); // Finaliza o carregamento
      });
  }, []);

  // Função para calcular os totais por pessoa
  const calcularTotaisPorPessoa = (transacoes) => {
    const totaisPorPessoa = {};

    transacoes.forEach((transacao) => {
      const pessoaID = transacao.Pessoa;

      // Inicializa os totais para a pessoa, se necessário
      if (!totaisPorPessoa[pessoaID]) {
        totaisPorPessoa[pessoaID] = {
          totalReceitas: 0,
          totalDespesas: 0,
          saldoLiquido: 0,
        };
      }

      // Atualiza os totais
      if (transacao.Tipo === "receita") {
        totaisPorPessoa[pessoaID].totalReceitas += transacao.Valor;
      } else if (transacao.Tipo === "despesa") {
        totaisPorPessoa[pessoaID].totalDespesas += transacao.Valor;
      }

      totaisPorPessoa[pessoaID].saldoLiquido =
        totaisPorPessoa[pessoaID].totalReceitas - totaisPorPessoa[pessoaID].totalDespesas;
    });

    setTotais(totaisPorPessoa);
  };

  // Função para cadastrar uma pessoa
  const handleCadastrarPessoa = (dados) => {
    fetch("http://localhost:8080/pessoas", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(dados),
    })
      .then((response) => response.json())
      .then((data) => {
        setPessoas((prev) => [...prev, data]);
      })
      .catch((error) => {
        console.error("Erro ao cadastrar pessoa:", error);
      });
  };

  // Função para cadastrar uma transação
  const handleCadastrarTransacao = (dados) => {
    fetch("http://localhost:8080/transacoes", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(dados),
    })
      .then(() => {
        // Recarrega as transações após o cadastro
        fetch("http://localhost:8080/transacoes/listar")
          .then((response) => response.json())
          .then((data) => {
            setTransacoes(data || []); // Garante que transacoes seja um array, mesmo se data for null
            calcularTotaisPorPessoa(data || []); // Recalcula os totais por pessoa após cadastrar uma transação
          });
      })
      .catch((error) => {
        console.error("Erro ao cadastrar transação:", error);
      });
  };

  // Exibe um indicador de carregamento enquanto os dados são carregados
  if (loading) {
    return <div>Carregando...</div>;
  }

  return (
    <div className="App">
      <h1>Controle de Gastos Residenciais</h1>
      <FormPessoa onSubmit={handleCadastrarPessoa} />
      <FormTransacao pessoas={pessoas} onSubmit={handleCadastrarTransacao} />
      <Totais totais={totais} pessoas={pessoas} />
    </div>
  );
}

export default App;