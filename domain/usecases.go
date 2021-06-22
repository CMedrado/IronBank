package domain

import (
	store_account "github.com/CMedrado/DesafioStone/store/account"
	store_token "github.com/CMedrado/DesafioStone/store/token"
	store_transfer "github.com/CMedrado/DesafioStone/store/transfer"
)

type AccountUseCase interface {
	CreateAccount(name string, cpf string, secret string, balance int) (int, error)
	GetBalance(id int) (int, error)
	GetAccounts() []store_account.Account
	SearchAccount(id int) store_account.Account
	UpdateBalance(accountOrigin store_account.Account, accountDestination store_account.Account)
	GetAccountCPF(cpf string) store_account.Account
	GetAccount() []store_account.Account
	SearchAccountCPF(cpf string) store_account.Account
}

type TransferUseCase interface {
	GetTransfers(token string) ([]store_transfer.Transfer, error)
	CreateTransfers(token string, accountDestinationID int, amount int) (error, int)
}

type LoginUseCase interface {
	AuthenticatedLogin(cpf, secret string) (error, string)
	GetTokenID(id int) store_token.Token
}
