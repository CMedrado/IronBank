package account

import (
	"github.com/CMedrado/DesafioStone/domain"
	store_account "github.com/CMedrado/DesafioStone/store/account"
)

type UseCase struct {
	StoredAccount Repository
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
	newAccount := store_account.Account{ID: id, Name: name, CPF: cpf, Secret: secretHash, Balance: balance, CreatedAt: domain.CreatedAt()}
	auc.StoredAccount.CreateAccount(newAccount)
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
func (auc *UseCase) GetAccounts() []store_account.Account {
	accounts := auc.StoredAccount.GetAccounts()
	var account []store_account.Account

	for _, a := range accounts {
		account = append(account, a)
	}

	return account
}

// SearchAccount returns the account via the received ID
func (auc UseCase) SearchAccount(id int) store_account.Account {
	accounts := auc.StoredAccount.GetAccounts()
	account := store_account.Account{}

	for _, a := range accounts {
		if a.ID == id {
			account = a
		}
	}

	return account
}

func (auc UseCase) UpdateBalance(accountOrigin store_account.Account, accountDestination store_account.Account) {
	auc.StoredAccount.UpdateBalances(accountOrigin, accountDestination)
}

func (auc *UseCase) GetAccountCPF(cpf string) store_account.Account {
	return auc.StoredAccount.GetAccountCPF(cpf)
}

func (auc UseCase) GetAccount() map[string]store_account.Account {
	return auc.StoredAccount.GetAccounts()
}
