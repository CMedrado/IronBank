package domain

import "github.com/CMedrado/DesafioStone/store"

type MethodsDomain interface {
	GetAccounts() []store.Account
	GetBalance(int) (uint, error)
	CreateAccount(string, string, string, uint) (int, error)
	AuthenticatedLogin(string, string) (error, string)
	GetTransfers(string) ([]store.Transfer, error)
	CreateTransfers(string, int, uint) (error, int)
}
