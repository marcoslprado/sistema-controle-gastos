import React from "react";

function Totais({ totais, pessoas }) {
  return (
    <div className="Totais">
      <h2>Totais por Pessoa</h2>
      {Object.keys(totais).map((pessoaID) => {
        const pessoa = pessoas.find((p) => p.ID === parseInt(pessoaID));
        const { totalReceitas, totalDespesas, saldoLiquido } = totais[pessoaID];

        return (
          <div key={pessoaID}>
            <h3>{pessoa ? pessoa.Nome : "Pessoa Desconhecida"}</h3>
            <p>Total Receitas: R$ {totalReceitas.toFixed(2)}</p>
            <p>Total Despesas: R$ {totalDespesas.toFixed(2)}</p>
            <p>Saldo Líquido: R$ {saldoLiquido.toFixed(2)}</p>
          </div>
        );
      })}
    </div>
  );
}

export default Totais;