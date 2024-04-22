package account

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/CMedrado/DesafioStone/pkg/domain"
)

// GetBalance requests the salary for the Story by sending the ID
func (auc UseCase) GetBalance(id string) (int, error) {
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return 0, fmt.Errorf("given the UUID is incorrect: %w", err)
	}

	account, err := auc.StoredAccount.ReturnAccountID(idUUID)
	if err != nil {
		return 0, domain.ErrSelect
	}

	err = domain.CheckExistID(account)
	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}
