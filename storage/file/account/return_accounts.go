package account

func (a StoredAccount) ReturnAccounts() []Account {
	return a.accounts
}
