package authentication

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/storage/file/token"
)

//func ChangeTokenDomain(tokenDomain domain.Token) token.Token {
//	tokenStorage := token.Token{ID: tokenDomain.ID,Token: tokenDomain.Token}
//	return tokenStorage
//}

func ChangeTokenStorage(tokenStorage token.Token) domain.Token {
	tokenDomain := domain.Token{ID: tokenStorage.ID, Token: tokenStorage.Token}
	return tokenDomain
}
