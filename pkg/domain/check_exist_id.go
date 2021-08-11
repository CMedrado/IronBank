package domain

import "github.com/CMedrado/DesafioStone/pkg/domain/entities"

// CheckExistID checks if the id exists and returns nil if not, it returns an error
func CheckExistID(accountOrigin entities.Account) error {
	if (accountOrigin == entities.Account{}) {
		return ErrInvalidID
	}
	return nil
}
