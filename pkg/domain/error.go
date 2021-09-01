package domain

import (
	"errors"
)

var (
	ErrInvalidCPF   = errors.New("given cpf is invalid")
	ErrInvalidToken = errors.New("given token is invalid")
	ErrInvalidID    = errors.New("given id is invalid")
	ErrParse        = errors.New("given the UUID is incorrect")
	ErrInsert       = errors.New("unable to insert")
	ErrSelect       = errors.New("unable to select")
)
