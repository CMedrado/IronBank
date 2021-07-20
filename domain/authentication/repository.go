package authentication

import (
	"github.com/CMedrado/DesafioStone/storage/postgre/token"
)

type Repository interface {
	SaveToken(token token.Token) error
	ReturnTokens() ([]token.Token, error)
}
