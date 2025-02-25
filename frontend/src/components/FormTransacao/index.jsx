import React, { useState } from "react";

function FormTransacao({ pessoas, onSubmit }) {
  const [descricao, setDescricao] = useState("");
  const [valor, setValor] = useState("");
  const [tipo, setTipo] = useState("despesa");
  const [pessoaId, setPessoaId] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit({
      descricao,
      valor: parseFloat(valor),
      tipo,
      pessoa: parseInt(pessoaId),
    });
    setDescricao("");
    setValor("");
    setPessoaId("");
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Cadastrar Transação</h2>
      <input
        type="text"
        placeholder="Descrição"
        value={descricao}
        onChange={(e) => setDescricao(e.target.value)}
        required
      />
      <input
        type="number"
        placeholder="Valor"
        value={valor}
        onChange={(e) => setValor(e.target.value)}
        required
      />
      <select value={tipo} onChange={(e) => setTipo(e.target.value)}>
        <option value="despesa">Despesa</option>
        <option value="receita">Receita</option>
      </select>
      <select
        value={pessoaId}
        onChange={(e) => setPessoaId(e.target.value)}
        required
      >
        <option value="">Selecione uma pessoa</option>
        {pessoas.map((pessoa) => (
          <option key={pessoa.ID} value={pessoa.ID}>
            {pessoa.Nome}
          </option>
        ))}
      </select>
      <button type="submit">Cadastrar</button>
    </form>
  );
}

export default FormTransacao;