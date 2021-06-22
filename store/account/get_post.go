package account

func (a *StoredAccount) CreateAccount(account Account) {
	accountStorage[account.CPF] = account
}

func (a StoredAccount) UpdateBalances(person1, person2 Account) {
	accountStorage[person1.CPF] = person1
	accountStorage[person2.CPF] = person2
}

func (a StoredAccount) GetAccounts() map[string]Account {
	return accountStorage
}

func (a StoredAccount) GetAccountCPF(cpf string) Account {
	return accountStorage[cpf]
}
