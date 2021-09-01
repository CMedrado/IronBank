package domain

import (
	"errors"
)

var (
	ErrInvalidCPF   = errors.New("given cpf is invalid")
	ErrInvalidToken = errors.New("given token is invalid")
	ErrInvalidID    = errors.New("given id is invalid")
	ErrParse        = errors.New("given the UUID is incorrect")
	ErrInsert       = errors.New("there was an error trying the insert command in the database")
	ErrSelect       = errors.New("there was an error trying the select command in the database")
)
