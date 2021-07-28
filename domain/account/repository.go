package account

import (
	"github.com/CMedrado/DesafioStone/storage/postgre/accounts"
	"github.com/google/uuid"
)

type Repository interface {
	SaveAccount(account accounts.Account) error
	ReturnAccounts() ([]accounts.Account, error)
	ChangeBalance(person1, person2 accounts.Account) error
	ReturnAccountID(id uuid.UUID) (accounts.Account, error)
	ReturnAccountCPF(cpf string) (accounts.Account, error)
}
