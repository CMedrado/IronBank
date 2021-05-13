package domain

import (
	"errors"
	"github.com/CMedrado/DesafioStone/store"
)

var (
	ErrInvalidAccount = errors.New("given account is invalid")
	ErrInvalidSecret  = errors.New("given secret is invalid")
	ErrInvalidCPF     = errors.New("given cpf is invalid")
	ErrWithoutBalance = errors.New("account without balance")
	ErrInvalidToken   = errors.New("given token is invalid")
	ErrInvalidID      = errors.New("given id is invalid")
)

func CheckedError(cpf string) error {

	if len(cpf) != 11 && len(cpf) != 14 {
		return ErrInvalidCPF
	}

	if len(cpf) == 14 {
		if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
			return nil
		} else {
			return ErrInvalidCPF
		}
	}
	return nil
}

func CheckAccount(account store.Account) error {
	if (account == store.Account{}) {
		return ErrInvalidAccount
	}
	return nil
}

func CheckBalance(person1 store.Account, amount int) error {
	if person1.Balance < amount {
		return ErrWithoutBalance
	}
	return nil
}

func CheckLogin(account store.Account, newlogin store.Login) error {

	if account.CPF != newlogin.CPF {
		return ErrInvalidCPF
	}
	if account.Secret != newlogin.Secret {
		return ErrInvalidSecret
	}
	return nil
}

//func CheckTransfer(transfer store.Transfer) error {
//	if (transfer[id] == store.Transfer{}) {
//		return ErrInvalidID
//	}
//	return nil
//}

func CheckID(token int, accountOriginID int, tokenStore store.Token) error {
	if tokenStore.AccountOriginID != accountOriginID {
		return ErrInvalidID
	} else if tokenStore.Token != token || tokenStore.Token == 0 {
		return ErrInvalidToken
	}
	return nil
}
