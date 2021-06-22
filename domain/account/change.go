package account

import (
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/CMedrado/DesafioStone/storage/file/account"
)

func ChangeAccountDomain(accountDomain domain.Account) account.Account {
	accountStorage := account.Account{ID: accountDomain.ID, Name: accountDomain.Name, CPF: accountDomain.CPF, Secret: accountDomain.Secret, Balance: accountDomain.Balance, CreatedAt: accountDomain.CreatedAt}
	return accountStorage
}

func ChangeAccountStorage(accountStorage account.Account) domain.Account {
	accountDomain := domain.Account{ID: accountStorage.ID, Name: accountStorage.Name, CPF: accountStorage.CPF, Secret: accountStorage.Secret, Balance: accountStorage.Balance, CreatedAt: accountStorage.CreatedAt}
	return accountDomain
}
