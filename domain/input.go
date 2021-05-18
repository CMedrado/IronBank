package domain

import "github.com/CMedrado/DesafioStone/store"

type MethodsDomain interface {
	GetAccounts() []store.Account
	GetBalance(int) (uint, error)
	CreateAccount(string, string, string, uint) (int, error)
	AuthenticatedLogin(string, string) (error, string)
	GetTransfers(string) ([]store.Transfer, error)
	MakeTransfers(string, int, uint) (error, int)
}

//type MethodsStore interface {
//	TransferredAccount(string, store.Account)
//	TransferredBalance(string) store.Account
//	TransferredAccounts() map[string]store.Account
//	//CheckLogin(string) store.Account
//}
