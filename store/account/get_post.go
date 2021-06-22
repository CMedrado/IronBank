package account

import "encoding/json"

func (a *StoredAccount) CreateAccount(account Account) {
	a.accounts = append(a.accounts, account)
	a.dataBase.Seek(0, 0)
	json.NewEncoder(a.dataBase).Encode(a.accounts)
}

func (a *StoredAccount) UpdateBalances(person1, person2 Account) {
	accountOrigin := a.accounts.Find(person1.Name)
	accountDestination := a.accounts.Find(person2.Name)

	accountOrigin.Balance = person1.Balance
	accountDestination.Balance = person2.Balance
}

func (a StoredAccount) GetAccounts() []Account {
	return a.accounts
}
