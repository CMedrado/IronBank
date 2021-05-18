package domain

import (
	"encoding/base64"
	"github.com/CMedrado/DesafioStone/store"
	"strconv"
)

func (auc AccountUsecase) AuthenticatedLogin(cpf, secret string) (error, string) {
	secretHash := Hash(secret)

	err := CheckedError(cpf)
	cpf = CpfReplace(cpf)
	if err != nil {
		return err, ""
	}

	newLogin := store.Login{cpf, secretHash}
	account := auc.Store.CheckLogin(cpf)

	err = CheckLogin(account, newLogin)
	if err != nil {
		return err, ""
	}

	id := auc.Store.TransferredAccounts()
	now := CreatedAt()
	token := now + ":" + strconv.Itoa(id[cpf].ID)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	auc.Token.CreatedToken(id[cpf].ID, encoded)

	return nil, encoded
}
