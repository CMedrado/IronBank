package domain

import "github.com/CMedrado/DesafioStone/store"

type StorageMethods interface {
	GetAccounts() []store.Account
	GetBalance(int) int
	CreatedAccount(string, string, string)
}
