package domain

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func CheckAccountExistence(account entities.Account) error {
	if (account != entities.Account{}) {
		return ErrAccountExists
	}
	return nil
}
