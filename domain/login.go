package domain

import (
	"github.com/CMedrado/DesafioStone/store"
	"math/rand"
)

func AuthenticatedLogin(cpf, secret string) (bool, error, int) {
	newlogin := store.Login{cpf, secret, 0}
	account := store.StoredAccount{}.CheckLogin(cpf)
	if account.CPF != newlogin.CPF {
		return false, ErrInvalidCPF, 0
	}
	if account.Secret != newlogin.Secret {
		return false, ErrInvalidSecret, 0
	}
	token := rand.Intn(10000000)
	storelogin := store.StoredLogin{}
	storelogin.CreatedLogin(token, cpf, secret)
	return true, nil, token
}
