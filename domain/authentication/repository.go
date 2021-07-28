package authentication

import (
	"github.com/CMedrado/DesafioStone/storage/postgre/token"
	"github.com/google/uuid"
)

type Repository interface {
	SaveToken(token token.Token) error
	ReturnTokens() ([]token.Token, error)
	ReturnTokenID(id uuid.UUID) (token.Token, error)
}
