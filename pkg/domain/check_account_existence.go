package domain

import (
	account2 "github.com/CMedrado/DesafioStone/pkg/domain/account"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func CheckAccountExistence(account entities.Account) error {
	if (account != entities.Account{}) {
		return account2.ErrAccountExists
	}
	return nil
}
