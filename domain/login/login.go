package login

import (
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/domain"
	store_login "github.com/CMedrado/DesafioStone/store/login"
	"strconv"
)

type UseCase struct {
	AccountUseCase domain.AccountUseCase
	StoredToken    Repository
}

// AuthenticatedLogin authenticates the account and returns a token
func (auc UseCase) AuthenticatedLogin(cpf, secret string) (error, string) {
	secretHash := domain.CreateHash(secret)

	err := domain.CheckCPF(cpf)
	cpf = domain.CpfReplace(cpf)
	if err != nil {
		return domain.ErrLogin, ""
	}

	newLogin := store_login.Login{CPF: cpf, Secret: secretHash}
	account := auc.AccountUseCase.GetAccountCPF(cpf)

	err = domain.CheckLogin(account, newLogin)
	if err != nil {
		return domain.ErrLogin, ""
	}

	id := auc.AccountUseCase.SearchAccountCPF(cpf)
	now := domain.CreatedAt()
	token := now + ":" + strconv.Itoa(id.ID)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	auc.StoredToken.PostToken(id.ID, encoded)

	return nil, encoded
}

func (uc UseCase) GetTokenID(id int) store_login.Token {
	return uc.StoredToken.GetTokenID(id)
}
