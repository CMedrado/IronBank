package account

import "github.com/CMedrado/DesafioStone/pkg/domain/entities"

func (auc UseCase) UpdateBalance(accountOrigin entities.Account, accountDestination entities.Account) error {
	err := auc.StoredAccount.ChangeBalance(accountOrigin, accountDestination)
	if err != nil {
		return ErrUpdate
	}
	return nil
}
