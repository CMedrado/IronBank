package authentication

import "errors"

var (
	ErrLogin         = errors.New("given secret or CPF are incorrect")
	ErrInvalidSecret = errors.New("given secret is invalid")
)
