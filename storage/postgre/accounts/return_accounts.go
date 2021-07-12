package accounts

import (
	"context"
)

func (a *Storage) ReturnAccounts() ([]Account, error) {
	statement := `SELECT * FROM accounts`
	rows, err := a.pool.Query(context.Background(), statement)
	if err != nil {
		return []Account{}, err
	}
	defer rows.Close()
	var account Account
	var accounts []Account
	for rows.Next() {
		rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		accounts = append(accounts, account)
	}
	return accounts, nil
}
