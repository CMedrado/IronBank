package domain

import "github.com/CMedrado/DesafioStone/store"

type MethodsDomain interface {
	GetAccounts() []store.Account
	GetBalance(int) int
	CreatedAccount(string, string, string)
}

type MethodsStore interface {
	TransferredAccount(string, store.Account)
	TransferredBalance(string) store.Account
	TransferredAccounts() map[string]store.Account
	//CheckLogin(string) store.Account
}
