package store

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

func (a StoredAccount) ReturnCPF(cpf string) int {
	return accountStorage[cpf].ID
}

func (a *StoredToken) PostToken(id int, token string) {
	accountToken[id] = Token{Token: token}
}

func (a *StoredToken) GetTokenID(id int) Token {
	return accountToken[id]
}

func (a StoredToken) ReturnToken(id int) string {
	return accountToken[id].Token
}

func (a *StoredTransferAccountID) GetTransfers(accountOriginID int) map[int]Transfer {
	return accountTransferAccountID[accountOriginID].accountTransferID
}

func (a *StoredTransferID) PostTransferAccountID(transfer Transfer) StoredTransferID {
	accountTransferID[transfer.ID] = transfer
	return StoredTransferID{accountTransferID: accountTransferID}
}

func (a *StoredTransferAccountID) PostTransferID(transfer Transfer, id int) {
	storedTransfer := StoredTransferID{}
	transferAccount := storedTransfer.PostTransferAccountID(transfer)
	accountTransferAccountID[id] = transferAccount
}