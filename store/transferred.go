package store

func (a *StoredAccount) TransferredAccount(conta Account) {
	accountStorage[conta.CPF] = conta
}

func (a StoredAccount) TransferredBalance(cpf string) Account {
	return accountStorage[cpf]
}

func (a StoredAccount) TransferredAccounts() map[string]Account {
	return accountStorage
}

func (a StoredAccount) CheckLogin(cpf string) Account {
	return accountStorage[cpf]
}
