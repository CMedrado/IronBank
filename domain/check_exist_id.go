package domain

// CheckExistID checks if the id exists and returns nil if not, it returns an error
func CheckExistID(accountOrigin Account) error {
	if (accountOrigin == Account{}) {
		return ErrInvalidID
	}
	return nil
}
