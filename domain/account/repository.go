package account

import "github.com/CMedrado/DesafioStone/store"

type Repository interface {
	GetAccountCPF(cpf string) store.Account
	CreateAccount(account store.Account)
	GetAccounts() map[string]store.Account
	UpdateBalances(person1, person2 store.Account)
	ReturnCPF(cpf string) int
}
