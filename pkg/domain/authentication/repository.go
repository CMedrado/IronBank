package authentication

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

type Repository interface {
	SaveToken(token entries.Token) error
	ReturnTokenID(id uuid.UUID) (entries.Token, error)
}
