package account

import (
	"github.com/CMedrado/DesafioStone/storage/postgre/accounts"
)

type Repository interface {
	SaveAccount(account accounts.Account) error
	ReturnAccounts() ([]accounts.Account, error)
	ChangeBalance(person1, person2 accounts.Account) error
}
