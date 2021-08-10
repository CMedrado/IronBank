package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/pkg/domain/entries"
	"github.com/google/uuid"
)

func (a *Storage) ReturnAccountID(id uuid.UUID) (entries.Account, error) {
	var account entries.Account
	statement := `SELECT * FROM accounts WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return entries.Account{}, err
	}
	return account, nil
}
