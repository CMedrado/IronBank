package Domain

import (
	"github.com/CMedrado/DesafioStone/Store"
	"time"
)


func (a *Store.ArmazenamentoDeContas) CriarConta(name string, cpf string, secret string) {
	id := 5
	created_at := time.Now().Format("02/01/2006 03:03:05")
	contaNova := Store.Conta{id, name, cpf, secret, 0, created_at}
	a.armazenamento[id] = contaNova
}

func (a ArmazenamentoDeContas) MostrarSaldo(id int) int {
	conta := a.armazenamento[id]
	return conta.Balance
}

func InicializaConta() *ArmazenamentoDeContas {
	return &ArmazenamentoDeContas{map[int]Store.Conta{}}
}

func (a *ArmazenamentoDeContas) ObterContas() []Store.Conta {
	var contas []Store.Conta
	return contas
}
