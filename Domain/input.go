package Domain

import "github.com/CMedrado/DesafioStone/Store"

type MetodosDeArmazenamento interface {
	Store.ObterContas() []Store.Conta
	MostrarSaldo(int) int
	CriarConta(string, string, string)
}
