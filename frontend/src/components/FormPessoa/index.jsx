import React, { useState } from "react";

function FormPessoa({ onSubmit }) {
  const [nome, setNome] = useState("");
  const [idade, setIdade] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit({ nome, idade: parseInt(idade) });
    setNome("");
    setIdade("");
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Cadastrar Pessoa</h2>
      <input
        type="text"
        placeholder="Nome"
        value={nome}
        onChange={(e) => setNome(e.target.value)}
        required
      />
      <input
        type="number"
        placeholder="Idade"
        value={idade}
        onChange={(e) => setIdade(e.target.value)}
        required
      />
      <button type="submit">Cadastrar</button>
    </form>
  );
}

export default FormPessoa;