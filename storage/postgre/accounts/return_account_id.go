package accounts

import (
	"context"
	"github.com/google/uuid"
)

func (a *Storage) ReturnAccountID(id uuid.UUID) (Account, error) {
	var account Account
	statement := `SELECT accounts WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&account)
	if err != nil {
		return Account{}, err
	}
	return account, nil
}
