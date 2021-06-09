package account

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/store"
)

type UseCase struct {
	StoredAccount *store.StoredAccount
}

//CreateAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
func (auc *UseCase) CreateAccount(name string, cpf string, secret string, balance int) (int, error) {
	err := domain.CheckCPF(cpf)
	if err != nil {
		return 0, err
	}
	err = domain.CheckBalance(balance)
	if err != nil {
		return 0, err
	}
	id := domain.Random()
	secretHash := domain.CreateHash(secret)
	cpf = domain.CpfReplace(cpf)
	newAccount := store.Account{ID: id, Name: name, CPF: cpf, Secret: secretHash, Balance: balance, CreatedAt: domain.CreatedAt()}
	auc.StoredAccount.PostAccount(newAccount)
	return id, err
}

//GetBalance requests the salary for the Story by sending the ID
func (auc *UseCase) GetBalance(id int) (int, error) {
	account := auc.SearchAccount(id)
	err := domain.CheckExistID(account)

	if err != nil {
		return 0, err
	}

	return account.Balance, nil

}

//GetAccounts returns all API accounts
func (auc *UseCase) GetAccounts() []store.Account {
	accounts := auc.StoredAccount.GetAccounts()
	var account []store.Account

	for _, a := range accounts {
		account = append(account, a)
	}

	return account
}

// SearchAccount returns the account via the received ID
func (auc UseCase) SearchAccount(id int) store.Account {
	accounts := auc.StoredAccount.GetAccounts()
	account := store.Account{}

	for _, a := range accounts {
		if a.ID == id {
			account = a
		}
	}

	return account
}
