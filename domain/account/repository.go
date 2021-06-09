package account

import "github.com/CMedrado/DesafioStone/store"

type Repository interface {
	GetAccounts() []store.Account
	GetBalance(int) (int, error)
	CreateAccount(string, string, string, int) (int, error)
}
