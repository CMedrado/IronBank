package domain

import (
	"errors"
	"github.com/CMedrado/DesafioStone/store"
)

var (
	ErrInvalidConta   = errors.New("given account is invalid")
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

func CheckConta(account store.Account) error {
	if (account == store.Account{}) {
		return ErrInvalidConta
	}
	return nil
}

func CheckBalance(person1 store.Account, amount int) error {
	if person1.Balance < amount {
		return ErrWithoutBalance
	}
	return nil
}

func CheckLogin(login store.Login) error {
	if (login == store.Login{}) {
		return ErrInvalidToken
	}
	return nil
}

func CheckID(transfer store.Transfer) error {
	if (transfer == store.Transfer{}) {
		return ErrInvalidID
	}
	return nil
}
