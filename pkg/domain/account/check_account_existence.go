package account

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
)

func CheckAccountExistence(account entries.Account) error {
	if (account != entries.Account{}) {
		return domain2.ErrAccountExists
	}
	return nil
}
