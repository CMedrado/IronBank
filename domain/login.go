package domain

import (
	"crypto/sha256"
	"github.com/CMedrado/DesafioStone/store"
)

func (auc AccountUsecase) AuthenticatedLogin(cpf, secret string) (error, int) {
	secretHash := sha256.Sum256([]byte(secret))
	newLogin := store.Login{cpf, secretHash}
	account := auc.Store.CheckLogin(cpf)

	err := CheckLogin(account, newLogin)
	if err != nil {
		return err, 0
	}

	token := Random()
	id := auc.Store.TransferredAccounts()
	auc.Token.CreatedToken(id[cpf].ID, token)

	return nil, token
}
