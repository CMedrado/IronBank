package domain

import (
	"errors"
	"github.com/CMedrado/DesafioStone/store"
)

var (
	errInvalidSecret  = errors.New("given secret is invalid")
	errInvalidCPF     = errors.New("given cpf is invalid")
	errWithoutBalance = errors.New("account without balance")
	errInvalidToken   = errors.New("given token is invalid")
	errInvalidID      = errors.New("given id is invalid")
	errInvalidAmount  = errors.New("given amount is invalid")
)

func CheckedError(cpf string) error {

	if len(cpf) != 11 && len(cpf) != 14 {
		return errInvalidCPF
	}

	if len(cpf) == 14 {
		if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
			return nil
		} else {
			return errInvalidCPF
		}
	}
	return nil
}

func CheckBalance(person1 store.Account, amount uint) error {
	if person1.Balance < amount {
		return errWithoutBalance
	}
	return nil
}

func CheckLogin(account store.Account, newlogin store.Login) error {

	if account.CPF != newlogin.CPF {
		return errInvalidCPF
	}
	if account.Secret != newlogin.Secret {
		return errInvalidSecret
	}
	return nil
}

func CheckToken(token string, tokens store.Token) error {
	if token != tokens.Token {
		return errInvalidToken
	}
	return nil
}

func CheckExistID(account store.Account) error {
	if (account == store.Account{}) {
		return errInvalidID
	}
	return nil
}

func CheckAmount(amount uint) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	return nil
}
