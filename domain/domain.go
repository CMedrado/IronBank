package domain

import (
	"github.com/CMedrado/DesafioStone/store"
	"time"
)

func CriarConta(name string, cpf string, secret string) {
	id := 5
	created_at := time.Now().Format("02/01/2006 03:03:05")
	contaNova := store.Conta{id, name, cpf, secret, 0, created_at}
	store.TransferenciaDoArmazenamento(id, contaNova)
}

func (a ArmazenamentoDeContas) MostrarSaldo(id int) int {
	conta := a.armazenamento[id]
	return conta.Balance
}

func InicializaConta() *ArmazenamentoDeContas {
	return &ArmazenamentoDeContas{map[int]store.Conta{}}
}

func (a *ArmazenamentoDeContas) ObterContas() []Store.Conta {
	var contas []store.Conta
	return contas
}
