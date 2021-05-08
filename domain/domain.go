package domain

import (
	"errors"
)

var (
	ErrInvalidSecret = errors.New("given secret is invalid")
	ErrInvalidCPF    = errors.New("given cpf is invalid")
)

//func AuthenticatedLogin(cpf, secret string) (bool, error, int) {
//	newlogin := store.Login{cpf,secret}
//	account := store.StoredAccount{}.CheckLogin(cpf)
//	if (account.CPF != newlogin.CPF) {
//		return false,ErrInvalidCPF
//	}
//	if (account.Secret != newlogin.Secret) {
//		return false, ErrInvalidSecret
//	}
//	token := rand.Intn(10000000)
//	return true,nil, token
//}

//func InicializaConta() *ArmazenamentoDeContas {
//	return &ArmazenamentoDeContas{map[int]store.Account{}}
//}
