package transfer

import "errors"

var (
	ErrInvalidCredential = errors.New("given the authorization header type is not basic")
)
