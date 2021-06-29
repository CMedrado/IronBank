package login

import "github.com/CMedrado/DesafioStone/store"

type Repository interface {
	ReturnToken(id int) string
	PostToken(id int, token string)
	GetTokenID(id int) store.Token
}
