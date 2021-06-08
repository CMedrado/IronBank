package domain

import "github.com/CMedrado/DesafioStone/store"

type AccountRepository interface {
	GetAccounts() []store.Account
	GetBalance(int) (int, error)
	CreateAccount(string, string, string, int) (int, error)
}

type LoginRepository interface {
	AuthenticatedLogin(string, string) (error, string)
}

type TransferRepository interface {
	GetTransfers(string) ([]store.Transfer, error)
	CreateTransfers(string, int, int) (error, int)
}

