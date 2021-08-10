package accounts

import (
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
)

func ChangeAccountDomain(accountDomain entries.Account) Account {
	accountStorage := Account{ID: accountDomain.ID, Name: accountDomain.Name, CPF: accountDomain.CPF, Secret: accountDomain.Secret, Balance: accountDomain.Balance, CreatedAt: accountDomain.CreatedAt}
	return accountStorage
}
