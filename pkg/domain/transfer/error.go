package transfer

import "errors"

var (
	ErrWithoutBalance       = errors.New("given account without balance")
	ErrInvalidAmount        = errors.New("given amount is invalid")
	ErrSameAccount          = errors.New("given account is the same as the account destination")
	ErrInvalidDestinationID = errors.New("given account destination id is invalid")
)
