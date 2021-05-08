package store

func (a *StoredAccount) TransferredAccount(id int, conta Account) {
	a.storage[id] = conta
}

func (a StoredAccount) TransferredBalance(id int) Account {
	conta := a.storage[id]
	return conta
}

func (a StoredAccount) TransferredAccounts() map[int]Account {
	accounts := a.storage
	return accounts
}
