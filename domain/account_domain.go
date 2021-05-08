package domain

import (
	"github.com/CMedrado/DesafioStone/store"
	"math/rand"
	"time"
)

//CreatedAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
func CreatedAccount(name string, cpf string, secret string) (int, error) {
	id := rand.Intn(10000000)
	created_at := time.Now().Format("02/01/2006 03:03:05")
	newAccount := store.Account{id, name, cpf, secret, 0, created_at}
	store.StoredAccount.TransferredAccount(store.StoredAccount{}, id, newAccount)
	return id, CheckedError(cpf)
}

//GetBalance requests the salary for the Story by sending the ID
func GetBalance(id int) int {
	conta := store.StoredAccount{}.TransferredBalance(id)
	return conta.Balance
}

//GetAccounts s
func GetAccounts() []store.Account {
	accounts := store.StoredAccount{}.TransferredAccounts()
	var account []store.Account

	for _, a := range accounts {
		account = append(account, a)
	}

	return account
}
