package domain

import "github.com/CMedrado/DesafioStone/store"

type MethodsDomain interface {
	GetAccounts() []store.Account
	GetBalance(string) (uint, error)
	CreateAccount(string, string, string, uint) (int, error)
	AuthenticatedLogin(string, string) (error, int)
	GetTransfers(int, int) ([]store.Transfer, error)
	MakeTransfers(int, int, int, uint) (error, int)
}

//type MethodsStore interface {
//	TransferredAccount(string, store.Account)
//	TransferredBalance(string) store.Account
//	TransferredAccounts() map[string]store.Account
//	//CheckLogin(string) store.Account
//}
