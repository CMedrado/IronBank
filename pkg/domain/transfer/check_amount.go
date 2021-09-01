package transfer

// CheckAmount checks if the amount is valid and returns nil if not, it returns an error
func CheckAmount(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	return nil
}
