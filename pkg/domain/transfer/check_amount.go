package transfer

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
)

// CheckAmount checks if the amount is valid and returns nil if not, it returns an error
func CheckAmount(amount int) error {
	if amount <= 0 {
		return domain2.ErrInvalidAmount
	}
	return nil
}
