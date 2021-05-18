package domain

import (
	"github.com/CMedrado/DesafioStone/store"
)

type AccountUsecase struct {
	Store    *store.StoredAccount
	Login    *store.StoredLogin
	Token    *store.StoredToken
	Transfer *store.StoredTransferTwo
}

//CreateAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
func (auc *AccountUsecase) CreateAccount(name string, cpf string, secret string, balance uint) (int, error) {
	err := CheckedError(cpf)
	if err != nil {
		return 0, err
	} else {
		id := Random()
		secretHash := Hash(secret)
		cpf = CpfReplace(cpf)
		newAccount := store.Account{id, name, cpf, secretHash, balance, CreatedAt()}
		auc.Store.TransferredAccount(newAccount)
		return id, err
	}
}

//GetBalance requests the salary for the Story by sending the ID
func (auc *AccountUsecase) GetBalance(id int) (uint, error) {
	account, err := auc.SearchID(id)

	if err != nil {
		return 0, err
	}

	return account.Balance, nil

}

//GetAccounts s
func (auc *AccountUsecase) GetAccounts() []store.Account {
	accounts := auc.Store.TransferredAccounts()
	var account []store.Account

	for _, a := range accounts {
		account = append(account, a)
	}

	return account
}
