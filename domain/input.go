package domain

import "github.com/CMedrado/DesafioStone/store"

type MethodsDomain interface {
	GetAccounts() ([]store.Account, error)
	GetBalance(int) (int, error)
	CreatedAccount(string, string, string, int) (int, error)
	AuthenticatedLogin(cpf, secret string) (error, int)
	GetTransfers(accountOriginID int, token int) ([]store.Transfer, error)
	MakeTransfers(accountOriginID int, token int, accountDestinationID int, amount int) error
}

type MethodsStore interface {
	TransferredAccount(string, store.Account)
	TransferredBalance(string) store.Account
	TransferredAccounts() map[string]store.Account
	//CheckLogin(string) store.Account
}
