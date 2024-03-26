package account

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
)

type Repository interface {
	SaveAccount(account entities.Account) error
	ReturnAccounts() ([]entities.Account, error)
	ChangeBalance(accountOrigin, accountDestination entities.Account) error
	ReturnAccountID(id uuid.UUID) (entities.Account, error)
	ReturnAccountCPF(cpf string) (entities.Account, error)
}
