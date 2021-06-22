package account

import (
	store_account "github.com/CMedrado/DesafioStone/store/account"
)

type Repository interface {
	GetAccountCPF(cpf string) store_account.Account
	CreateAccount(account store_account.Account)
	GetAccounts() []store_account.Account
	UpdateBalances(person1, person2 store_account.Account)
}
