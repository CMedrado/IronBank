package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
	"github.com/google/uuid"
)

func (a *Storage) ReturnAccountID(id uuid.UUID) (domain.Account, error) {
	var account domain.Account
	statement := `SELECT * FROM accounts WHERE id=$1`
	err := a.pool.QueryRow(context.Background(), statement, id).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}
