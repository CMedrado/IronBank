package authentication

import (
	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

type Repository interface {
	SaveToken(token entities.Token) error
	ReturnTokenID(id uuid.UUID) (entities.Token, error)
	ReturnAccountCPF(cpf string) (entities.Account, error)
}
