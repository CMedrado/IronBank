package domain

import (
	"crypto/sha256"
	"github.com/CMedrado/DesafioStone/store"
)

type AccountUsecase struct {
	Store    *store.StoredAccount
	Login    *store.StoredLogin
	Token    *store.StoredToken
	Transfer *store.StoredTransferTwo
}

//CreateAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
func (auc *AccountUsecase) CreateAccount(name string, cpf string, secret string, balance int) (int, error) {
	err := CheckedError(cpf)
	if err != nil {
		return 0, ErrInvalidCPF
	} else {
		id := Random()
		secretHash := sha256.Sum256([]byte(secret))
		cpf = CpfReplace(cpf)
		newAccount := store.Account{id, name, cpf, secretHash, 0, CreatedAt()}
		auc.Store.TransferredAccount(newAccount)
		return id, err
	}
}

//GetBalance requests the salary for the Story by sending the ID
func (auc *AccountUsecase) GetBalance(cpf string) (int, error) {
	cpf = CpfReplace(cpf)
	conta := auc.Store.TransferredBalance(cpf)
	err := CheckAccount(conta)
	if err != nil {
		return 0, err
	}
	err = CheckedError(cpf)
	if err != nil {
		return 0, err
	}
	return conta.Balance, nil

}

//GetAccounts s
func (auc *AccountUsecase) GetAccounts() []store.Account {
	accounts := auc.Store.TransferredAccounts()
	var account []store.Account

	for _, a := range accounts {
		account = append(account, a)
	}

	return account
}

////CreateAccount to receive Name, CPF and Secret and set up the account, creating ID and Created_at
//func CreateAccount(name string, cpf string, secret string) (int, error) {
//	err := CheckedError(cpf)
//	if err != nil {
//		return 0, ErrInvalidCPF
//	} else {
//		id := Random()
//		created_at := CreatedAt()
//		newAccount := store.Account{id, name, cpf, secret, 0, created_at}
//		storeMethod := store.StoredAccount{}
//		storeMethod.TransferredAccount(newAccount)
//		return id, err
//	}
//}
//
////GetBalance requests the salary for the Story by sending the ID
//func GetBalance(cpf string) (int, error) {
//	storeMethod := store.StoredAccount{}
//	conta := storeMethod.TransferredBalance(cpf)
//	err := CheckAccount(conta)
//	if err != nil {
//		return 0, err
//	} else {
//		return conta.Balance, nil
//	}
//}
//
////GetAccounts s
//func GetAccounts() []store.Account {
//	storeMethod := store.StoredAccount{}
//	accounts := storeMethod.TransferredAccounts()
//	var account []store.Account
//
//	for _, a := range accounts {
//		account = append(account, a)
//	}
//
//	return account
//}
