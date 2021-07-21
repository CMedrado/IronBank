package account

import "github.com/CMedrado/DesafioStone/domain"

// CheckBalance checks if the amount is valid and returns nil if not, it returns an error
func CheckBalance(amount int) error {
	if amount < 0 {
		return domain.ErrInvalidAmount
	}
	return nil
}
