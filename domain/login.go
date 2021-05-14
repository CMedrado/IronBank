package domain

import (
	"github.com/CMedrado/DesafioStone/store"
)

func (auc AccountUsecase) AuthenticatedLogin(cpf, secret string) (error, int) {
	secretHash := Hash(secret)

	err := CheckedError(cpf)
	if err != nil {
		return err, 0
	}

	cpf = CpfReplace(cpf)
	newLogin := store.Login{cpf, secretHash}
	account := auc.Store.CheckLogin(cpf)

	err = CheckLogin(account, newLogin)
	if err != nil {
		return err, 0
	}

	token := Random()
	id := auc.Store.TransferredAccounts()
	auc.Token.CreatedToken(id[cpf].ID, token)

	return nil, token
}
