package domain

import "github.com/CMedrado/DesafioStone/store"

type MetodosDeArmazenamento interface {
	GetAccounts() []store.Account
	GetBalance(int) int
	CreatedAccount(string, string, string)
}
