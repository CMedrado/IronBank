package authentication

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
	"github.com/google/uuid"
)

type Repository interface {
	SaveToken(token entities.Token) error
	ReturnTokenID(id uuid.UUID) (entities.Token, error)
}
