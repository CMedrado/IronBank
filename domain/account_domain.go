package domain

import (
	"github.com/CMedrado/DesafioStone/store"
	"math/rand"
	"time"
)

//CreatedAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
func CreatedAccount(name string, cpf string, secret string) (int, error) {
	err := CheckedError(cpf)
	if err != nil {
		return 0, ErrInvalidCPF
	} else {
		id := rand.Intn(10000000)
		created_at := time.Now().Format("02/01/2006 03:03:05")
		newAccount := store.Account{id, name, cpf, secret, 0, created_at}
		storeMethod := store.StoredAccount{}
		storeMethod.TransferredAccount(newAccount)
		return id, err
	}
}

//GetBalance requests the salary for the Story by sending the ID
func GetBalance(cpf string) (int, error) {
	storeMethod := store.StoredAccount{}
	conta := storeMethod.TransferredBalance(cpf)
	err := CheckConta(conta)
	if err != nil {
		return 0, err
	}
	return conta.Balance, nil
}

//GetAccounts s
func GetAccounts() []store.Account {
	storeMethod := store.StoredAccount{}
	accounts := storeMethod.TransferredAccounts()
	var account []store.Account

	for _, a := range accounts {
		account = append(account, a)
	}

	return account
}
