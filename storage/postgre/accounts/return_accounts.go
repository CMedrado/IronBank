package accounts

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
)

func (a *Storage) ReturnAccounts() ([]domain.Account, error) {
	statement := `SELECT * FROM accounts`
	rows, err := a.pool.Query(context.Background(), statement)
	if err != nil {
		return []domain.Account{}, err
	}
	defer rows.Close()
	var account domain.Account
	var accounts []domain.Account
	for rows.Next() {
		rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		accounts = append(accounts, account)
	}
	return accounts, nil
}
