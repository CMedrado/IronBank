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

func (a *StoredLogin) CreatedLogin(token int, cpf string, secret string) {
	accountLogin[token] = Login{cpf, secret, token}
}

func (a StoredLogin) GetLogin(token int) Login {
	return accountLogin[token]
}

func (a StoredTransfer) GetTransfers() map[int]map[int]Transfer {
	return accountTransfer
}

func (a StoredAccount) UpdateBalance(person1, person2 Account) {
	accountStorage[person1.CPF] = person1
	accountStorage[person2.CPF] = person2
}

func (a StoredTransfer) CreatedTransfer(transfer Transfer) {
	accountTransfer[transfer.ID][transfer.AccountOriginID] = transfer
}
