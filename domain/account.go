package domain

import (
	"github.com/CMedrado/DesafioStone/store"
)

type AccountUseCase struct {
	Store    *store.StoredAccount
	Login    *store.StoredLogin
	Token    *store.StoredToken
	Transfer *store.StoredTransferAccountID
}

//CreateAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
func (auc *AccountUseCase) CreateAccount(name string, cpf string, secret string, balance uint) (int, error) {
	err := CheckCPF(cpf)
	if err != nil {
		return 0, err
	} else {
		id := Random()
		secretHash := CreateHash(secret)
		cpf = CpfReplace(cpf)
		newAccount := store.Account{ID: id, Name: name, CPF: cpf, Secret: secretHash, Balance: balance, CreatedAt: CreatedAt()}
		auc.Store.PostAccount(newAccount)
		return id, err
	}
}

//GetBalance requests the salary for the Story by sending the ID
func (auc *AccountUseCase) GetBalance(id int) (uint, error) {
	account, err := auc.SearchAccount(id)

	if err != nil {
		return 0, err
	}

	return account.Balance, nil

}

//GetAccounts returns all API accounts
func (auc *AccountUseCase) GetAccounts() []store.Account {
	accounts := auc.Store.GetAccounts()
	var account []store.Account

	for _, a := range accounts {
		account = append(account, a)
	}

	return account
}
