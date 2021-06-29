package account

import (
	store_account "github.com/CMedrado/DesafioStone/store/account"
)

type Repository interface {
	CreateAccount(account store_account.Account)
	GetAccounts() []store_account.Account
	UpdateBalances(person1, person2 store_account.Account)
}
