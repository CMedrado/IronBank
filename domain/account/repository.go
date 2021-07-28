package account

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

type Repository interface {
	SaveAccount(account domain.Account) error
	ReturnAccounts() ([]domain.Account, error)
	ChangeBalance(person1, person2 domain.Account) error
	ReturnAccountID(id uuid.UUID) (domain.Account, error)
	ReturnAccountCPF(cpf string) (domain.Account, error)
}
