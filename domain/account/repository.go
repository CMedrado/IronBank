package account

import "github.com/CMedrado/DesafioStone/store"

type Repository interface {
	GetAccountCPF(string) store.Account
	CreateAccount(account store.Account) (int, error)
	GetAccounts() map[string]store.Account
	UpdateBalance(person1, person2 store.Account)
	ReturnCPF(cpf string) int
}
