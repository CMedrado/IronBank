package store

func (a *StoredAccount) PostAccount(account Account) {
	accountStorage[account.CPF] = account
}

func (a StoredAccount) GetBalance(cpf string) Account {
	return accountStorage[cpf]
}

func (a StoredAccount) GetAccounts() map[string]Account {
	return accountStorage
}

func (a StoredAccount) GetAccountCPF(cpf string) Account {
	return accountStorage[cpf]
}

func (a StoredLogin) GetLogin(token int) Login {
	return accountLogin[token]
}

func (a *StoredToken) PostToken(id int, token string) {
	accountToken[id] = Token{token}
}

func (a *StoredToken) GetTokenID(id int) Token {
	return accountToken[id]
}

func (a *StoredTransferID) GetTransfers(accountOriginID int) map[int]Transfer {
	return accountTransferID[accountOriginID].accountTransferAccountID
}

func (a StoredAccount) UpdateBalance(person1, person2 Account) {
	accountStorage[person1.CPF] = person1
	accountStorage[person2.CPF] = person2
}

func (a *StoredTransferAccountID) PostTransferAccountID(transfer Transfer) StoredTransferAccountID {
	accountTransferAccountID[transfer.ID] = transfer
	return StoredTransferAccountID{accountTransferAccountID}
}

func (a *StoredTransferID) PostTransferID(transfer Transfer, id int) {
	storedTransfer := StoredTransferAccountID{}
	transferAccount := storedTransfer.PostTransferAccountID(transfer)
	accountTransferID[id] = transferAccount
}

func (a StoredAccount) ReturnCPF(cpf string) int {
	return accountStorage[cpf].ID
}

func (a StoredToken) ReturnToken(id int) string {
	return accountToken[id].Token
}
