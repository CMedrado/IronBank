package authentication

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/storage/postgre/token"
)

//func ChangeTokenDomain(tokenDomain domain.Token) token.Token {
//	tokenStorage := token.Token{ID: tokenDomain.ID,Token: tokenDomain.Token}
//	return tokenStorage
//}

func ChangeTokenStorage(tokenStorage token.Token) domain.Token {
	tokenDomain := domain.Token{ID: tokenStorage.ID, IdAccount: tokenStorage.IdAccount, CreatedAt: tokenStorage.CreatedAt}
	return tokenDomain
}
