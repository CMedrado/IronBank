package authentication

import (
	"github.com/CMedrado/DesafioStone/storage/postgre/token"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	SaveToken(idToken uuid.UUID, id uuid.UUID, time time.Time) error
	ReturnTokens() ([]token.Token, error)
}
