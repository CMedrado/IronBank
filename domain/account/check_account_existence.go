package account

import "github.com/CMedrado/DesafioStone/domain"

func CheckAccountExistence(account domain.Account) error {
	if (account != domain.Account{}) {
		return domain.ErrAccountExists
	}
	return nil
}
