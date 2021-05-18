package store

func (a *StoredAccount) TransferredAccount(account Account) {
	accountStorage[account.CPF] = account
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

func (a StoredLogin) GetLogin(token int) Login {
	return accountLogin[token]
}

func (a *StoredToken) CreatedToken(id int, token string) {
	accountToken[id] = Token{token}
}

func (a *StoredToken) GetTokenID(id int) Token {
	return accountToken[id]
}

func (a *StoredTransferTwo) GetTransfers(s int) map[int]Transfer {
	return accountTransferTwo[s].accountTransfer
}

func (a StoredAccount) UpdateBalance(person1, person2 Account) {
	accountStorage[person1.CPF] = person1
	accountStorage[person2.CPF] = person2
}

func (a *StoredTransfer) CreatedTransfer(transfer Transfer) StoredTransfer {
	accountTransfer[transfer.ID] = transfer
	return StoredTransfer{accountTransfer: accountTransfer}
}

func (a *StoredTransferTwo) CreatedTransferTwo(transfer Transfer, id int) {
	storedTransfer := StoredTransfer{}
	s := storedTransfer.CreatedTransfer(transfer)
	accountTransferTwo[id] = s
}

func (a StoredAccount) VoltaCPF(cpf string) int {
	return accountStorage[cpf].ID
}

func (a StoredToken) VoltaToken(id int) string {
	return accountToken[id].Token
}
