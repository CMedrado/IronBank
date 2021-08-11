package account

import (
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/entities"
)

func CheckAccountExistence(account entities.Account) error {
	if (account != entities.Account{}) {
		return domain2.ErrAccountExists
	}
	return nil
}
