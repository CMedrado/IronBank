package login

import (
	store_token "github.com/CMedrado/DesafioStone/store/token"
)

type Repository interface {
	PostToken(id int, token string)
	GetTokens() []store_token.Token
}
