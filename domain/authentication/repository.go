package authentication

import (
	"github.com/CMedrado/DesafioStone/storage/postgre/token"
	"github.com/google/uuid"
)

type Repository interface {
	SaveToken(idToken uuid.UUID, id uuid.UUID, time string) error
	ReturnTokens() ([]token.Token, error)
}
