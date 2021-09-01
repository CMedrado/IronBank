package transfer

// CheckAccountBalance checks if the account has a balance and returns nil if not, it returns an error
func CheckAccountBalance(person1 int, amount int) error {
	if person1 < amount {
		return ErrWithoutBalance
	}
	return nil
}
