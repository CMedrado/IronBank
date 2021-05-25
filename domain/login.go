package domain

import (
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/store"
	"strconv"
)

// AuthenticatedLogin authenticates the account and returns a token
func (auc AccountUseCase) AuthenticatedLogin(cpf, secret string) (error, string) {
	secretHash := CreateHash(secret)

	err := CheckCPF(cpf)
	cpf = CpfReplace(cpf)
	if err != nil {
		return ErrLogin, ""
	}

	newLogin := store.Login{CPF: cpf, Secret: secretHash}
	account := auc.Store.GetAccountCPF(cpf)

	err = CheckLogin(account, newLogin)
	if err != nil {
		return ErrLogin, ""
	}

	id := auc.Store.GetAccounts()
	now := CreatedAt()
	token := now + ":" + strconv.Itoa(id[cpf].ID)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	auc.Token.PostToken(id[cpf].ID, encoded)

	return nil, encoded
}
