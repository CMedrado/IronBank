package store

func (a *StoredAccount) TransferredAccount(cpf string, conta Account) {
	a.storage[cpf] = conta
}

func (a StoredAccount) TransferredBalance(cpf string) Account {
	conta := a.storage[cpf]
	return conta
}

func (a StoredAccount) TransferredAccounts() map[string]Account {
	accounts := a.storage
	return accounts
}

func (a StoredAccount) CheckLogin(cpf string) Account {
	account := a.storage[cpf]

	return account
}
