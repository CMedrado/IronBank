package domain

import "github.com/CMedrado/DesafioStone/store"

type AccountUsecase interface {
	CreateAccount(name string, cpf string, secret string, balance int) (int, error)
	GetBalance(id int) (int, error)
	GetAccounts() []store.Account
	SearchAccount(id int) store.Account
	UpdateBalance(accountOrigin store.Account, accountDestination store.Account)
}

type TransferUseCase interface {
	GetTransfers(token string) ([]store.Transfer, error)
	CreateTransfers(token string, accountDestinationID int, amount int) (error, int)
}
