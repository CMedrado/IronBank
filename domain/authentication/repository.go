package authentication

import (
	"github.com/CMedrado/DesafioStone/storage/file/token"
	"github.com/google/uuid"
)

type Repository interface {
	SaveToken(id uuid.UUID, token string)
	ReturnTokens() []token.Token
}
