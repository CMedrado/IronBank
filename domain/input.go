package domain

import "github.com/CMedrado/DesafioStone/store"

type MetodosDeArmazenamento interface {
	ObterContas() []store.Conta
	MostrarSaldo(int) int
	CriarConta(string, string, string)
}
