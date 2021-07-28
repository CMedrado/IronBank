package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

func (a *Storage) ReturnAccountID(id uuid.UUID) (domain.Account, error) {
	var account domain.Account
	statement := `SELECT accounts WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&account)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}
