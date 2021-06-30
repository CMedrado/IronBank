package transfer

import "github.com/CMedrado/DesafioStone/domain"

// CheckAccountBalance checks if the account has a balance and returns nil if not, it returns an error
func CheckAccountBalance(person1 int, amount int) error {
	if person1 < amount {
		return domain.ErrWithoutBalance
	}
	return nil
}
