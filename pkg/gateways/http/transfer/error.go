package transfer

import "errors"

var (
	ErrInvalidCredential = errors.New("given the credential is not basic")
)
