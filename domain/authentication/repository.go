package authentication

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

type Repository interface {
	SaveToken(token domain.Token) error
	ReturnTokenID(id uuid.UUID) (domain.Token, error)
}
