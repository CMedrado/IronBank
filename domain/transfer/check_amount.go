package transfer

import "github.com/CMedrado/DesafioStone/domain"

// CheckAmount checks if the amount is valid and returns nil if not, it returns an error
func CheckAmount(amount int) error {
	if amount <= 0 {
		return domain.ErrInvalidAmount
	}
	return nil
}
