package domain

import (
	"errors"
)

var (
	ErrInvalidSecret        = errors.New("given secret is invalid")
	ErrInvalidCPF           = errors.New("given cpf is invalid")
	ErrWithoutBalance       = errors.New("given account without balance")
	ErrInvalidToken         = errors.New("given token is invalid")
	ErrInvalidID            = errors.New("given id is invalid")
	ErrInvalidAmount        = errors.New("given amount is invalid")
	ErrInvalidDestinationID = errors.New("given account destination id is invalid")
	ErrSameAccount          = errors.New("given account is the same as the account destination")
	ErrBalanceAbsent        = errors.New("given the balance amount is invalid")
	ErrLogin                = errors.New("given secret or CPF are incorrect")
	ErrAccountExists        = errors.New("given cpf is already used")
	ErrParse                = errors.New("given the UUID is incorrect")
	ErrInsert               = errors.New("unable to insert")
	ErrUpdate               = errors.New("unable to update")
)
