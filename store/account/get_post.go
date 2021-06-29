package account

import "encoding/json"

func (a *StoredAccount) CreateAccount(account Account) {
	a.accounts = append(a.accounts, account)
	a.dataBase.Seek(0, 0)
	json.NewEncoder(a.dataBase).Encode(a.accounts)
}

func (a StoredAccount) UpdateBalances(person1, person2 Account) {
	accountStorage[person1.CPF] = person1
	accountStorage[person2.CPF] = person2
}

func (a StoredAccount) GetAccounts() []Account {
	return a.accounts
}

func (a StoredAccount) GetAccountCPF(cpf string) Account {
	return accountStorage[cpf]
}
