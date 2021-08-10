package domain

import "github.com/CMedrado/DesafioStone/pkg/domain/entries"

// CheckExistID checks if the id exists and returns nil if not, it returns an error
func CheckExistID(accountOrigin entries.Account) error {
	if (accountOrigin == entries.Account{}) {
		return ErrInvalidID
	}
	return nil
}
