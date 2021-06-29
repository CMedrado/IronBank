package domain

import "github.com/CMedrado/DesafioStone/store"

type AccountUseCase interface {
	CreateAccount(name string, cpf string, secret string, balance int) (int, error)
	GetBalance(id int) (int, error)
	GetAccounts() []store.Account
	SearchAccount(id int) store.Account
	UpdateBalance(accountOrigin store.Account, accountDestination store.Account)
	ReturnCPF(cpf string) int
	GetAccountCPF(cpf string) store.Account
	GetAccount() map[string]store.Account
}

type TransferUseCase interface {
	GetTransfers(token string) ([]store.Transfer, error)
	CreateTransfers(token string, accountDestinationID int, amount int) (error, int)
}

type LoginUseCase interface {
	AuthenticatedLogin(cpf, secret string) (error, string)
	ReturnToken(id int) string
	GetTokenID(id int) store.Token
}
