package account

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

type Repository interface {
	SaveAccount(account entries.Account) error
	ReturnAccounts() ([]entries.Account, error)
	ChangeBalance(person1, person2 entries.Account) error
	ReturnAccountID(id uuid.UUID) (entries.Account, error)
	ReturnAccountCPF(cpf string) (entries.Account, error)
}
