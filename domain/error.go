package domain

import (
	"errors"
	"github.com/CMedrado/DesafioStone/store"
)

var (
	errInvalidSecret        = errors.New("given secret is invalid")
	errInvalidCPF           = errors.New("given cpf is invalid")
	errWithoutBalance       = errors.New("given account without balance")
	errInvalidToken         = errors.New("given token is invalid")
	errInvalidID            = errors.New("given id is invalid")
	errInvalidAmount        = errors.New("given amount is invalid")
	errInvalidDestinationID = errors.New("given account destination id is invalid")
	errSameAccount          = errors.New("given account is the same as the account destination")
	errBalanceAbsent        = errors.New("given the balance amount is invalid")
	ErrLogin                = errors.New("given secret or CPF are incorrect")
)

// CheckCPF checks if the cpf exists and returns nil if not, it returns an error
func CheckCPF(cpf string) error {

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

// CheckAccountBalance checks if the account has a balance and returns nil if not, it returns an error
func CheckAccountBalance(person1 int, amount int) error {
	if person1 < amount {
		return errWithoutBalance
	}
	return nil
}

// CheckLogin Checks if the cpf and secret ar correct and returns nil if not, it returns an error
func CheckLogin(account store.Account, newLogin store.Login) error {

	if account.CPF != newLogin.CPF {
		return errInvalidCPF
	}
	if account.Secret != newLogin.Secret {
		return errInvalidSecret
	}
	return nil
}

// CheckToken checks if the token is correct and returns nil if not, it returns an error
func CheckToken(token string, tokens store.Token) error {
	if token != tokens.Token {
		return errInvalidToken
	}
	return nil
}

// CheckExistID checks if the id exists and returns nil if not, it returns an error
func CheckExistID(account store.Account) error {
	if (account == store.Account{}) {
		return errInvalidID
	}
	return nil
}

// CheckAmount checks if the amount is valid and returns nil if not, it returns an error
func CheckAmount(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	return nil
}

// CheckCompareID Compare two IDs to see if they are the same and returns nil if not, it returns an error
func CheckCompareID(accountOriginID, accountDestinationID int) error {
	if accountOriginID == accountDestinationID {
		return errSameAccount
	}
	return nil
}

// CheckExistDestinationID checks if the destination id exists and returns nil if not, it returns an error
func CheckExistDestinationID(account store.Account) error {
	if (account == store.Account{}) {
		return errInvalidDestinationID
	}
	return nil
}

// CheckBalance checks if the balance exists and returns nil if not, it returns an error
func CheckBalance(balance int) error {
	if balance <= 0 {
		return errBalanceAbsent
	}
	return nil
}
