package account

import "io"

type Accounts []Account

type Account struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Secret    string `json:"secret"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}

type StoredAccount struct {
	dataBase io.ReadWriteSeeker
	accounts Accounts
}

func NewStoredAccount(dataBase io.ReadWriteSeeker) *StoredAccount {
	dataBase.Seek(0, 0)
	accounts, _ := NewAccount(dataBase)

	return &StoredAccount{dataBase: dataBase, accounts: accounts}
}
