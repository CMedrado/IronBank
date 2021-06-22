package account

import (
	"github.com/CMedrado/DesafioStone/storage/file/account"
)

type Repository interface {
	SaveAccount(account account.Account)
	ReturnAccounts() []account.Account
	ChangeBalances(person1, person2 account.Account)
}
