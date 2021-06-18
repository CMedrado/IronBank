package login

import (
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
	"strconv"
)

type UseCase struct {
	StoredAccount *store.StoredAccount
	StoredToken   *store.StoredToken
}

// AuthenticatedLogin authenticates the account and returns a token
func (auc UseCase) AuthenticatedLogin(cpf, secret string) (error, string) {
	secretHash := domain.CreateHash(secret)

	err := domain.CheckCPF(cpf)
	cpf = domain.CpfReplace(cpf)
	if err != nil {
		return domain.ErrLogin, ""
	}

	newLogin := store.Login{CPF: cpf, Secret: secretHash}
	account := auc.StoredAccount.GetAccountCPF(cpf)

	err = domain.CheckLogin(account, newLogin)
	if err != nil {
		return domain.ErrLogin, ""
	}

	id := auc.StoredAccount.GetAccounts()
	now := domain.CreatedAt()
	token := now + ":" + strconv.Itoa(id[cpf].ID)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	auc.StoredToken.PostToken(id[cpf].ID, encoded)

	return nil, encoded
}

func (uc UseCase) ReturnToken(id int) string {
	return uc.StoredToken.ReturnToken(id)
}

func (uc UseCase) GetTokenID(id int) store.Token {
	return uc.StoredToken.GetTokenID(id)
}
