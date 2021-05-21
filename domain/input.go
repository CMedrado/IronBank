package domain

import "github.com/CMedrado/DesafioStone/store"

type MethodsDomain interface {
	GetAccounts() []store.Account
	GetBalance(int) (int, error)
	CreateAccount(string, string, string, int) (int, error)
	AuthenticatedLogin(string, string) (error, string)
	GetTransfers(string) ([]store.Transfer, error)
	CreateTransfers(string, int, int) (error, int)
}
