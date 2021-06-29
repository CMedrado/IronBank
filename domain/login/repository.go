package login

import (
	store_login "github.com/CMedrado/DesafioStone/store/login"
)

type Repository interface {
	PostToken(id int, token string)
	GetTokenID(id int) store_login.Token
}
