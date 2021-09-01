package account

import "errors"

var (
	ErrAccountExists = errors.New("given cpf is already used")
	ErrBalanceAbsent = errors.New("given the balance amount is invalid")
	ErrUpdate        = errors.New("unable to update")
)
