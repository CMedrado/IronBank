package login

import (
	"github.com/CMedrado/DesafioStone/storage/file/token"
)

type Repository interface {
	SaveToken(id int, token string)
	ReturnTokens() []token.Token
}
